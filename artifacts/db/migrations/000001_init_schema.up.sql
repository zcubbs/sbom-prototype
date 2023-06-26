CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE artifacts
(
  id               BIGSERIAL PRIMARY KEY,
  artifact_name    VARCHAR(255) NOT NULL,
  artifact_type    VARCHAR(255) NOT NULL,
  artifact_version VARCHAR(255) NOT NULL,
  uuid             UUID         NOT NULL DEFAULT uuid_generate_v4(),
  created_at       TIMESTAMPTZ           DEFAULT current_timestamp,
  updated_at       TIMESTAMPTZ           DEFAULT current_timestamp,
  CONSTRAINT uc_artifacts_unique_info UNIQUE (artifact_name, artifact_type, artifact_version)
);

CREATE TABLE artifact_tags
(
  id          BIGSERIAL PRIMARY KEY,
  artifact_id BIGINT,
  tag_label   VARCHAR(255) NOT NULL,
  tag_value   VARCHAR(255) NOT NULL,
  FOREIGN KEY (artifact_id) REFERENCES artifacts (id),
  CONSTRAINT uc_artifact_tags UNIQUE (artifact_id, tag_label, tag_value)
);
