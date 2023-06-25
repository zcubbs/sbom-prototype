package grpc

import (
	"context"
	"fmt"
	"github.com/zcubbs/zlogger/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "zel/sbom-prototype/scanner/_gen/go/v1"
	"zel/sbom-prototype/scanner/internal/config"
)

type Scanner struct {
	log logger.Logger
}

func NewServer(l logger.Logger) *Scanner {
	return &Scanner{l}
}

func (s *Scanner) Scan(ctx context.Context, req *pb.ScanRequest) (*pb.ScanResponse, error) {
	s.log.Infof("Handle scan request for %s", req.Sbom)
	return &pb.ScanResponse{ReportId: "1"}, nil
}

func (s *Scanner) ScanResults(ctx context.Context, req *pb.ScanResultsRequest) (*pb.ScanResultsResponse, error) {
	s.log.Infof("Handle scanResults request for %s", req.ReportId)
	return &pb.ScanResultsResponse{
		Report: &pb.ScanReport{
			Vulnerabilities: []string{"none"},
		},
	}, nil
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
