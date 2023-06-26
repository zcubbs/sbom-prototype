## Table: artifacts

| Column           | Data Type    | Constraints                          | Description                                     |
|------------------|--------------|--------------------------------------|-------------------------------------------------|
| id               | BIGSERIAL    | PRIMARY KEY                          | Unique identifier for each artifact entry.      |
| uuid             | UUID         | NOT NULL, DEFAULT uuid_generate_v4() | Universally unique identifier for the artifact. |
| artifact_name    | VARCHAR(255) | NOT NULL                             | Name of the artifact.                           |
| artifact_type    | VARCHAR(255) | NOT NULL                             | Type of the artifact.                           |
| artifact_version | VARCHAR(255) | NOT NULL                             | Version of the artifact.                        |

Indexes:
- None

## Table: artifact_tags

| Column      | Data Type    | Constraints | Description                                    |
|-------------|--------------|-------------|------------------------------------------------|
| id          | BIGSERIAL    | PRIMARY KEY | Unique identifier for each artifact tag entry. |
| artifact_id | BIGINT       | NOT NULL    | Foreign key referencing the artifact.          |
| tag_label   | VARCHAR(255) | NOT NULL    | Label of the tag.                              |
| tag_value   | VARCHAR(255) | NOT NULL    | Value of the tag.                              |

Indexes:
- None
