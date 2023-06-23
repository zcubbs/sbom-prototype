package main

import "github.com/zcubbs/zlogger/pkg/logger"

func main() {
	SetupLogger(ZapLogger)

	img := "docker.io/library/alpine:latest"
	logger.L().Infof("scanning image: %s", img)
	sbom, scanReport, err := RunScan(img)
	if err != nil {
		logger.L().Fatalf("error: %s", err.Error())
	}

	logger.L().Infof("sbom:\n %s\nscanReport:\n%s", sbom, scanReport)
}

func init() {
	LoadConfig()
}
