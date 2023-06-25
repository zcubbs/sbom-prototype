package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zcubbs/zlogger/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	pb "zel/sbom-prototype/scanner/_gen/go/v1"
	"zel/sbom-prototype/scanner/internal/config"
	scannerGrpc "zel/sbom-prototype/scanner/internal/grpc"
	//pb "zel/sbom-prototype/scanner/proto/sbom/v1"
	"zel/sbom-prototype/scanner/sql"
)

func run(cfg config.Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Start gRPC server
	logger.L().Infof("Starting gRPC server on port %s", cfg.Grpc.Server.Port)
	go scannerGrpc.StartGrpcServer(cfg.Grpc.Server)

	// Register gRPC server Gateway endpoint
	grpcServPort := fmt.Sprintf(":%d", cfg.Grpc.Server.Port)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterScannerServiceHandlerFromEndpoint(ctx, mux, grpcServPort, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	logger.L().Infof("Starting server on port %d", cfg.HttpServer.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.HttpServer.Port), mux)
}

func main() {
	flag.Parse()

	cfg := config.Bootstrap()

	logger.SetupLogger(cfg.LoggerType)

	sql.MigrateDB(cfg.Database)

	if err := run(cfg); err != nil {
		logger.L().Fatal(err)
	}
}

//func main() {
//	logger.SetupLogger(logger.ZapLogger)
//
//	img := "docker.io/library/alpine:latest"
//	logger.L().Infof("scanning image: %s", img)
//	sbom, scanReport, err := RunScan(img)
//	if err != nil {
//		logger.L().Fatalf("error: %s", err.Error())
//	}
//
//	logger.L().Infof("sbom:\n %s\nscanReport:\n%s", sbom, scanReport)
//}
//
//func init() {
//	LoadConfig()
//}
