package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zcubbs/zlogger/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	pb "zel/sbom-prototype/scanner/proto/sbom/v1"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9000", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterScannerServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	logger.L().Infof("Starting HTTP server on port 8001")

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8001", mux)
}

func main() {
	flag.Parse()

	logger.SetupLogger(logger.ZapLogger)

	if err := run(); err != nil {
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
