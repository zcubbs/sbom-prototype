package application

import "zel/sbom-prototype/scanner/pkg/domain/scanner"

type Service interface {
	AddScan(scanner.Scan)
}

type service struct {
	repo scanner.Repository
}

func NewService(repo scanner.Repository) Service {
	return &service{repo: repo}
}

func (s *service) AddScan(scan scanner.Scan) {
	s.repo.AddScan(scan)
}
