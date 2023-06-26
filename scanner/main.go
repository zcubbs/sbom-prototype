package main

import (
	"context"
	"flag"
	"github.com/zcubbs/zlogger/pkg/logger"
	"zel/sbom-prototype/scanner/db"
	"zel/sbom-prototype/scanner/internal/config"
	scannerGrpc "zel/sbom-prototype/scanner/internal/grpc"
)

func main() {
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Bootstrap configuration
	cfg := config.Bootstrap()

	// Setup logger
	logger.SetupLogger(cfg.LoggerType)

	// Connect to database
	err := config.GetDbConnection(cfg.Database)
	if err != nil {
		logger.L().Fatal(err, "failed to connect to database")
	}

	// Migrate database
	db.MigrateDB(cfg.Database)

	// Open Nats connection
	err = config.ConnectToNats(cfg.Nats)
	if err != nil {
		logger.L().Fatal(err, "failed to connect to nats")
	}

	// Start telemetry service
	//defer config.ConfigTelemetry(context.Background())()

	// Start gRPC server
	logger.L().Infof("Starting gRPC server on port %s", cfg.Grpc.Server.Port)
	go scannerGrpc.StartGrpcServer(cfg.Grpc.Server)

	// Start gRPC gateway
	logger.L().Infof("Starting gRPC gateway on port %s", cfg.HttpServer.Port)
	err = scannerGrpc.StartGrpcGateway(cfg.Grpc.Server, cfg.HttpServer, ctx)
	if err != nil {
		logger.L().Fatal(err)
	}
}
