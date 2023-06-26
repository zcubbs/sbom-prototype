package scheduler

import (
	"context"
	"database/sql"
	"github.com/nats-io/nats.go"
	"github.com/zcubbs/zlogger/pkg/logger"
	"zel/sbom-prototype/scanner/internal/handler"
)

type Scheduler struct {
	natsConn *nats.Conn
	logger   logger.Logger
	DBConn   *sql.DB
	Ctx      context.Context
}

func New(natsConn *nats.Conn, l logger.Logger) *Scheduler {
	return &Scheduler{
		natsConn: natsConn,
		logger:   l,
	}
}

func (s *Scheduler) HandleScanJobsQueue() error {
	_, err := s.natsConn.Subscribe("hooks.gitlab.jobs", func(msg *nats.Msg) {
		s.logger.Infof("Received a message: %s\n", string(msg.Data))

		h := handler.New(s.logger, s.DBConn, s.Ctx)

		scan, err := h.GetScan(string(msg.Data))
		if err != nil {
			s.logger.Error(err)
		}

		response, err := h.RunScan(scan.Image)
		if err != nil {
			s.logger.Error(err)
		}

		s.logger.Infof("response: %+v\n", response)

	})

	if err != nil {
		return err
	}

	return nil
}
