package tasks

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/zcubbs/zlogger/pkg/logger"
)

type TaskProcessor interface {
	Start() error
	ProcessScanImageTask(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
	logger logger.Logger
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 1,
	})
	return &RedisTaskProcessor{server: server, store: store}
}

func (p *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeScanImage, p.ProcessScanImageTask)
	return p.server.Run(mux)
}
