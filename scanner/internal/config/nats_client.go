package config

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/zcubbs/zlogger/pkg/logger"
)

var NatsConn *nats.Conn

func ConnectToNats(cfg NatsConfig) error {
	logger.L().Info("Connecting to NATS server...")

	s := fmt.Sprintf("nats://%s:%d", cfg.Host, cfg.Port)

	nc, err := nats.Connect(s)
	if err != nil {
		return err
	}

	NatsConn = nc

	logger.L().Info("Connected to NATS server")

	return nil
}
