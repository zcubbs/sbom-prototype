package tasks

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/zcubbs/zlogger/pkg/logger"
)

type TaskDispatcher interface {
	DispatchScanImageTask(
		ctx context.Context,
		payload *ScanImagePayload,
		opts ...asynq.Option,
	) error
}

type RedisTaskDispatcher struct {
	client *asynq.Client
	logger logger.Logger
}

func NewRedisTaskDispatcher(redisOpt asynq.RedisClientOpt) TaskDispatcher {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDispatcher{client: client}
}
