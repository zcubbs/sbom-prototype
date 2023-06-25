package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"github.com/zcubbs/zlogger/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	pb "zel/sbom-prototype/scanner/_gen/go/v1"
	"zel/sbom-prototype/scanner/internal/config"
)

func StartGrpcGateway(grpcCfg config.GrpcServer, httpCfg config.HttpServerConfig, ctx context.Context) error {
	// Register gRPC server Gateway endpoint
	grpcServPort := fmt.Sprintf(":%d", grpcCfg.Port)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterScannerServiceHandlerFromEndpoint(ctx, mux, grpcServPort, opts)
	if err != nil {
		return err
	}

	// cors
	withCors := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(mux)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	logger.L().Infof("Starting server on port %d", httpCfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", httpCfg.Port), withCors)
}
