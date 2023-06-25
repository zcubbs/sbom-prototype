package grpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/zcubbs/zlogger/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
	pb "zel/sbom-prototype/scanner/_gen/go/v1"
	"zel/sbom-prototype/scanner/internal/config"
	"zel/sbom-prototype/scanner/internal/handler"
)

type Scanner struct {
	log logger.Logger
}

func NewServer(l logger.Logger) *Scanner {
	return &Scanner{l}
}

func (s *Scanner) AddScanImage(ctx context.Context, req *pb.AddScanImageRequest) (*pb.AddScanImageResponse, error) {
	s.log.Infof("Handle AddScanImage request for %s", req.Image)
	h := handler.New(s.log, config.DbConn, ctx)
	uid, err := h.ScheduleScan(req.Image)
	if err != nil {
		return nil, err
	}

	s.log.Infof("scheduled job id: %+v\n", uid)

	return &pb.AddScanImageResponse{
		JobId: uid,
	}, nil
}

func (s *Scanner) AddScanSbom(_ context.Context, req *pb.AddScanSbomRequest) (*pb.AddScanSbomResponse, error) {
	s.log.Infof("Handle AddScanSbom request for %s", req.Sbom)
	return &pb.AddScanSbomResponse{JobId: "1"}, nil
}

func (s *Scanner) GetScan(_ context.Context, req *pb.GetScanRequest) (*pb.GetScanResponse, error) {
	s.log.Infof("Handle scanResults request for %s", req.Uuid)
	return &pb.GetScanResponse{
		Scan: &pb.Scan{
			Uuid:            uuid.New().String(),
			Status:          "",
			CreatedAt:       "",
			UpdatedAt:       "",
			Image:           "test-image",
			ImageTag:        "",
			SbomId:          "",
			ArtifactId:      "",
			ArtifactName:    "",
			ArtifactVersion: "",
			Report:          nil,
		},
	}, nil
}

func (s *Scanner) GetScans(ctx context.Context, req *pb.GetScansRequest) (*pb.GetScansResponse, error) {
	s.log.Infof("Handle GetScans request req=%+v", req)

	h := handler.New(s.log, config.DbConn, ctx)

	response, err := h.GetScans(req.Limit, req.Page)
	if err != nil {
		return nil, err
	}

	parsedScans := make([]*pb.Scan, len(response.Scans))
	for i, scan := range response.Scans {
		parsedScans[i] = &pb.Scan{
			Uuid:            scan.Uuid,
			Status:          scan.Status,
			CreatedAt:       scan.CreatedAt.Format(time.RFC3339),
			UpdatedAt:       scan.UpdatedAt.Format(time.RFC3339),
			Image:           scan.Image,
			ImageTag:        scan.Image,
			SbomId:          scan.SbomID,
			ArtifactId:      scan.ArtifactID,
			ArtifactName:    scan.ArtifactName,
			ArtifactVersion: scan.ArtifactVersion,
			Report:          nil,
		}
	}

	return &pb.GetScansResponse{
		Scans: parsedScans,
		Pagination: &pb.Pagination{
			Count: response.Count,
			Pages: response.Pages,
		},
	}, nil
}

func (s *Scanner) RetryScan(context.Context, *pb.RetryScanRequest) (*pb.RetryScanResponse, error) {

	return &pb.RetryScanResponse{}, nil
}

func StartGrpcServer(cfg config.GrpcServer) {
	var s *grpc.Server
	// Load our TLS certificate and use grpc/credentials to create new transport credentials
	if cfg.TlsEnabled { // TLS is enabled
		logger.L().Infof("TLS is enabled, loading certificate from %s", cfg.TlsCertPem)
		sv, err := GenerateTLSServer(cfg)
		if err != nil {
			log.Fatal(err)
		}
		s = sv
	} else { // TLS is disabled
		s = grpc.NewServer()
	}

	// Register the grpc server
	ps := NewServer(logger.L())

	pb.RegisterScannerServiceServer(s, ps)

	// start the grpc server
	// TODO: add a flag-switch to enable/disable this feature
	reflection.Register(s)

	addr := fmt.Sprintf(":%d", cfg.Port)
	// Start the server
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listening on %s. ", addr)
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Successfully started gRPC server")
}

func GenerateTLSServer(cfg config.GrpcServer) (*grpc.Server, error) {
	cred, err := credentials.NewServerTLSFromFile(
		cfg.TlsCertPem,
		cfg.TlsCertKey,
	)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer(
		grpc.Creds(cred),
	)
	return s, nil
}
