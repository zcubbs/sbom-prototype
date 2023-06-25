package handler

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/zcubbs/zlogger/pkg/logger"
	"time"
	"zel/sbom-prototype/scanner/internal/runtime"
	"zel/sbom-prototype/scanner/internal/store"
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

func (h *Handler) ScheduleScan(image string) (scheduleUuid string, err error) {
	s := store.New(h.db)
	scan, err := s.CreateScan(h.ctx, store.CreateScanParams{
		Uuid:            uuid.New(),
		Image:           image,
		Sbom:            sql.NullString{},
		Status:          "pending",
		ArtifactID:      uuid.NullUUID{},
		ArtifactName:    sql.NullString{},
		ArtifactVersion: sql.NullString{},
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
	if err != nil {
		return "", err
	}

	h.logger.Infof("scan: %+v\n", scan)

	return scan.Uuid.String(), nil
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

type GetScanResponse struct {
	Uuid            string
	Image           string
	Sbom            string
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
	s := store.New(h.db)

	if limit == 0 {
		limit = 10
	}

	page++

	offset := (page - 1) * limit

	if limit < 1 || limit > 100 {
		limit = 10
	}

	scans, err := s.GetScans(h.ctx, store.GetScansParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var scansResponse []*GetScanResponse

	for _, scan := range scans {
		scansResponse = append(scansResponse, &GetScanResponse{
			Uuid:            scan.Uuid.String(),
			Image:           scan.Image,
			Sbom:            scan.Sbom.String,
			Status:          scan.Status,
			ArtifactID:      scan.ArtifactID.UUID.String(),
			ArtifactName:    scan.ArtifactName.String,
			ArtifactVersion: scan.ArtifactVersion.String,
			CreatedAt:       scan.CreatedAt,
			UpdatedAt:       scan.UpdatedAt,
		})
	}

	count, err := s.CountScans(h.ctx)
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
