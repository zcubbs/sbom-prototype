package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	app "zel/sbom-prototype/scanner/pkg/application/scanner"
	"zel/sbom-prototype/scanner/pkg/domain/scanner"
	grpc_scan "zel/sbom-prototype/scanner/pkg/infrastructure/delivery/grpc/proto/scan"
)

type Service struct {
	svc app.Service
}

func NewScanServerGrpc(gserver *grpc.Server, svc app.Service) {
	attenserver := &Service{svc: svc}

	grpc_scan.RegisterScanServiceServer(gserver, attenserver)
	reflection.Register(gserver)
}

func (s *Service) parseToGrpc(scan scanner.Scan) *grpc_scan.Scan {
	return &grpc_scan.Scan{
		Id:    scan.Id,
		Image: scan.Image,
	}
}

func (s *Service) parseToData(scan *grpc_scan.Scan) scanner.Scan {
	return scanner.Scan{
		Id:    scan.Id,
		Image: scan.Image,
	}
}

func (s *Service) Create(ctx context.Context, u *grpc_scan.CreateRequest) (*grpc_scan.CreateResponse, error) {
	// FIXME: implement errors
	a := s.parseToData(u.Scan)

	s.svc.AddScan(*a)

	return &grpc_scan.CreateResponse{}, nil
}
