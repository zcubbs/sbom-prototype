package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const TypeSbomImage = "scan:sbom"

type ScanSbomPayload struct {
	JobUUID uuid.UUID `json:"job_uuid"`
	Sbom    string    `json:"sbom"`
}

func (d *RedisTaskDispatcher) DispatchScanSbomTask(payload *ScanSbomPayload, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("json.Marshal failed: err=%v - %w", err, asynq.SkipRetry)
	}

	task := asynq.NewTask(TypeSbomImage, jsonPayload, opts...)
	info, err := d.client.Enqueue(task)
	if err != nil {
		return fmt.Errorf("could not enqueue task: %v - %w", err, asynq.SkipRetry)
	}

	d.logger.Infof("enqueued task: type=%s id=%s queue=%s payload=[% x]", task.Type(), info.ID, info.Queue, task.Payload())

	return nil
}
