package application

import (
	"context"
	"fmt"
)

func RunScan(image string) (sbomJson string, scanReportJson string, err error) {
	// Run image scan
	var img string

	if image == "" {
		img = "docker.io/library/alpine:latest"
	} else {
		img = image
	}

	r, err := NewRuntime(context.Background())
	if err != nil {
		panic(err)
	}

	sbom, err := r.GenerateSBOM(img)
	if err != nil {
		panic(err)
	}

	fmt.Printf("sbom: %+v\n", sbom)

	scan, err := r.ScanSBOM(sbom)
	if err != nil {
		panic(err)
	}

	fmt.Printf("scan results: %+v\n", scan)

	report, err := r.ParseVulnerabilityReport(scan)
	if err != nil {
		return
	}

	fmt.Printf("report: %+v\n", report)

	levels, fixes := r.ParseVulnerabilityForSeverityLevels(report)

	for level, count := range levels {
		fmt.Printf("Found %d %s vulnerabilities\n", count, level)
	}
	fmt.Printf("%d vulnerabilities have fixes available\n", fixes)

	return sbom, scan, nil
}
