CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE scan_jobs
(
  id               BIGSERIAL PRIMARY KEY,
  created_at       TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
  updated_at       TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
  sbom_uuid        UUID,
  artifact_uuid    UUID,
  artifact_name    VARCHAR(255) NOT NULL,
  artifact_version VARCHAR(255) NOT NULL,
  artifact_type    VARCHAR(255) NOT NULL,
  status           VARCHAR(255) NOT NULL,
  report           JSONB,
  job_log          TEXT
);

CREATE INDEX idx_scan_jobs_artifact_uuid ON scan_jobs (artifact_uuid);
CREATE INDEX idx_scan_jobs_sbom_uuid ON scan_jobs (sbom_uuid);
