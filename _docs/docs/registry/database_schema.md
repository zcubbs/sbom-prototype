## Table: sboms

| Column           | Data Type    | Constraints                          | Description                                     |
|------------------|--------------|--------------------------------------|-------------------------------------------------|
| id               | BIGSERIAL    | PRIMARY KEY                          | Unique identifier for each sbom entry.          |
| uuid             | UUID         | NOT NULL, DEFAULT uuid_generate_v4() | Universally unique identifier for the sbom.     |
| created_at       | TIMESTAMPTZ  | DEFAULT current_timestamp            | Timestamp indicating when the sbom was created. |
| manifest         | JSONB        |                                      | JSON object storing the sbom manifest.          |
| job_log          | TEXT         |                                      | Text field for storing job logs.                |
| status           | VARCHAR(255) | NOT NULL                             | Status of the sbom.                             |
| artifact_uuid    | UUID         | NOT NULL                             | Universally unique identifier for the artifact. |
| artifact_name    | VARCHAR(255) | NOT NULL                             | Name of the artifact.                           |
| artifact_version | VARCHAR(255) | NOT NULL                             | Version of the artifact.                        |
| artifact_type    | VARCHAR(255) | NOT NULL                             | Type of the artifact.                           |

Unique Constraint: uc_sbom_artifact_unique_info (artifact_uuid, artifact_name, artifact_version, artifact_type)

Indexes:
- idx_sbom_artifact_name (artifact_name)
- idx_sbom_uuid (uuid)
