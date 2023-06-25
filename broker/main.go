package main

import (
	"github.com/nats-io/nats-server/v2/server"
	"github.com/zcubbs/zlogger/pkg/logger"
	"time"
)

func main() {
	logger.SetupLogger(logger.ZapLogger)
	logger.L().Info("Starting NATS server...")
	opts := &server.Options{}
	opts.Port = 4222

	// Initialize new server with options
	ns, err := server.NewServer(opts)

	if err != nil {
		logger.L().Panic(err)
	}

	err = ns.SetDefaultSystemAccount()
	if err != nil {
		return
	}

	// Start the server via goroutine
	go ns.Start()
	logger.L().Infof("NATS server started - %s", ns.ClientURL())
	logger.L().Infof("Account: %v", ns.SystemAccount())

	// Wait for server to be ready for connections
	if !ns.ReadyForConnections(4 * time.Second) {
		logger.L().Panic("not ready for connection")
	}

	// Wait for server shutdown
	ns.WaitForShutdown()
}
