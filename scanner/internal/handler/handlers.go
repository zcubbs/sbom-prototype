package handler

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
	"github.com/zcubbs/zlogger/pkg/logger"
	"strconv"
	"time"
	db "zel/sbom-prototype/scanner/db/sqlc"
	"zel/sbom-prototype/scanner/internal/runtime"
)

type Handler struct {
	ctx    context.Context
	logger logger.Logger
	db     *sql.DB
}

func New(l logger.Logger, db *sql.DB, ctx context.Context) *Handler {
	return &Handler{
		logger: l,
		db:     db,
		ctx:    ctx,
	}
}

type RunScanResponse struct {
	Sbom   string
	Report string
	Image  string
}

func (h *Handler) ScheduleScan(image string) (id string, err error) {
	s := db.NewStore(h.db)
	scan, err := s.InsertScanJob(h.ctx, db.InsertScanJobParams{
		ArtifactUuid:    uuid.NullUUID{},
		ArtifactName:    "test",
		ArtifactVersion: "1.0.0",
		ArtifactType:    "image",
		Status:          "pending",
		Report:          pqtype.NullRawMessage{},
		JobLog:          sql.NullString{},
	})
	if err != nil {
		return "", err
	}

	h.logger.Infof("scan: %+v\n", scan)

	return strconv.FormatInt(scan.ID, 10), nil
}

func (h *Handler) RunScan(image string) (response *RunScanResponse, err error) {
	//Run image scan
	r, err := runtime.New(context.Background())
	if err != nil {
		h.logger.Error(err)
		return nil, err
	}

	sbom, err := r.GenerateSBOM(image)
	if err != nil {
		h.logger.Error(err)
		return nil, err
	}

	h.logger.Infof("sbom: %+v\n", sbom)

	scan, err := r.ScanSBOM(sbom)
	if err != nil {
		h.logger.Error(err)
		return nil, err
	}

	h.logger.Infof("scan results: %+v\n", scan)

	report, err := r.ParseVulnerabilityReport(scan)
	if err != nil {
		h.logger.Error(err)
		return nil, err
	}

	h.logger.Infof("report: %+v\n", report)

	levels, fixes := r.ParseVulnerabilityForSeverityLevels(report)

	for level, count := range levels {
		h.logger.Infof("Found %d %s vulnerabilities\n", count, level)
	}
	h.logger.Infof("%d vulnerabilities have fixes available\n", fixes)

	return &RunScanResponse{
		Sbom:   sbom,
		Report: scan,
		Image:  image,
	}, nil
}

func (h *Handler) GetScan(id int64) (response *db.ScanJob, err error) {
	s := db.NewStore(h.db)

	scan, err := s.GetScanJobByID(h.ctx, id)
	if err != nil {
		return nil, err
	}

	return &scan, nil
}

type GetScanResponse struct {
	Uuid            string
	Image           string
	ImageTag        string
	SbomID          string
	Status          string
	ArtifactID      string
	ArtifactName    string
	ArtifactVersion string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type GetScansResponse struct {
	Scans   []*GetScanResponse
	Count   int32
	Pages   int32
	Next    int32
	Prev    int32
	Current int32
}

func (h *Handler) GetScans(limit, page int32) (response *GetScansResponse, err error) {
	s := db.NewStore(h.db)

	if limit == 0 {
		limit = 10
	}

	page++

	offset := (page - 1) * limit

	if limit < 1 || limit > 100 {
		limit = 10
	}

	scans, err := s.GetScanJobsList(h.ctx, db.GetScanJobsListParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var scansResponse []*GetScanResponse

	for _, scan := range scans {
		scansResponse = append(scansResponse, &GetScanResponse{
			Uuid:            scan.ArtifactUuid.UUID.String(),
			Image:           scan.ArtifactName,
			ImageTag:        scan.ArtifactVersion,
			SbomID:          scan.SbomUuid.UUID.String(),
			Status:          scan.Status,
			ArtifactID:      scan.ArtifactUuid.UUID.String(),
			ArtifactName:    scan.ArtifactName,
			ArtifactVersion: scan.ArtifactVersion,
			CreatedAt:       time.Time{},
			UpdatedAt:       time.Time{},
		})
	}

	count, err := s.CountScanJobs(h.ctx)
	if err != nil {
		return nil, err
	}

	h.logger.Infof("limit: %d, page: %d, offset: %d, count: %d", limit, page, offset, count)

	return &GetScansResponse{
		Scans:   scansResponse,
		Count:   int32(count),
		Pages:   int32(count) / limit,
		Next:    page + 1,
		Prev:    page - 1,
		Current: page,
	}, nil
}
