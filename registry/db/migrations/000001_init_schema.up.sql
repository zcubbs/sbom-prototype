CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE sboms
(
  id               BIGSERIAL PRIMARY KEY,
  uuid             UUID         NOT NULL DEFAULT uuid_generate_v4(),
  created_at       TIMESTAMPTZ           DEFAULT current_timestamp,
  manifest         JSONB,
  job_log          TEXT,
  status           VARCHAR(255) NOT NULL,
  artifact_uuid    UUID         NOT NULL,
  artifact_name    VARCHAR(255) NOT NULL,
  artifact_version VARCHAR(255) NOT NULL,
  artifact_type    VARCHAR(255) NOT NULL,
  CONSTRAINT uc_sbom_artifact_unique_info UNIQUE (artifact_uuid, artifact_name, artifact_version, artifact_type)
);

CREATE INDEX idx_sbom_artifact_name ON sboms (artifact_name);
CREATE INDEX idx_sbom_uuid ON sboms (uuid);
