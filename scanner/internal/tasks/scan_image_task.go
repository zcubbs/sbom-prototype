package tasks

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const TypeScanImage = "scan:image"

type ScanImagePayload struct {
	JobUUID uuid.UUID `json:"job_uuid"`
	Image   string    `json:"image"`
}

func (d *RedisTaskDispatcher) DispatchScanImageTask(ctx context.Context, payload *ScanImagePayload, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("json.Marshal failed: err=%v - %w", err, asynq.SkipRetry)
	}

	task := asynq.NewTask(TypeScanImage, jsonPayload, opts...)
	info, err := d.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("could not enqueue task: %v - %w", err, asynq.SkipRetry)
	}

	d.logger.Infof("enqueued task: type=%s id=%s queue=%s payload=[% x]", task.Type(), info.ID, info.Queue, task.Payload())

	return nil
}

func (p *RedisTaskProcessor) ProcessScanImageTask(ctx context.Context, task *asynq.Task) error {
	var payload ScanImagePayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %w", asynq.SkipRetry)
	}

	scan, err := p.store.GetScan(ctx, payload.JobUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("scan not found: %w", asynq.SkipRetry)
		}
		return fmt.Errorf("could not get scan: %w", err)
	}

	// TODO: scan image
	p.logger.Infof("scan image: %s", payload.Image)

	return nil
}
