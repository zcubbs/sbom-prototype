## Table: scan_jobs

| Column           | Data Type    | Constraints               | Description                                                |
|------------------|--------------|---------------------------|------------------------------------------------------------|
| id               | BIGSERIAL    | PRIMARY KEY               | Unique identifier for each scan job entry.                 |
| created_at       | TIMESTAMPTZ  | DEFAULT current_timestamp | Timestamp indicating when the scan job was created.        |
| updated_at       | TIMESTAMPTZ  | DEFAULT current_timestamp | Timestamp indicating when the scan job was last updated.   |
| artifact_uuid    | UUID         | NOT NULL                  | Universally unique identifier for the associated artifact. |
| artifact_name    | VARCHAR(255) | NOT NULL                  | Name of the associated artifact.                           |
| artifact_version | VARCHAR(255) | NOT NULL                  | Version of the associated artifact.                        |
| artifact_type    | VARCHAR(255) | NOT NULL                  | Type of the associated artifact.                           |
| status           | VARCHAR(255) | NOT NULL                  | Status of the scan job.                                    |
| report           | JSONB        |                           | JSON object storing the scan job report.                   |
| job_log          | TEXT         |                           | Text field for storing job logs.                           |

Indexes:
- None
