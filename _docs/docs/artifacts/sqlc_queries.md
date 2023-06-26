### artifacts Table Queries

#### InsertArtifact
- Description: Insert a new artifact into the "artifacts" table.
- SQL:
    ```sql
    INSERT INTO artifacts (
      uuid,
      artifact_name,
      artifact_type,
      artifact_version
    ) VALUES (
      $1,
      $2,
      $3,
      $4
    ) RETURNING id;
    ```

#### GetArtifactByID
- Description: Retrieve an artifact from the "artifacts" table by ID.
- SQL:
    ```sql
    SELECT *
    FROM artifacts
    WHERE id = $1;
    ```

#### UpdateArtifact
- Description: Update an existing artifact in the "artifacts" table.
- SQL:
    ```sql
    UPDATE artifacts
    SET
      uuid = $1,
      artifact_name = $2,
      artifact_type = $3,
      artifact_version = $4
    WHERE id = $5;
    ```

#### DeleteArtifact
- Description: Delete an artifact from the "artifacts" table.
- SQL:
    ```sql
    DELETE FROM artifacts
    WHERE id = $1;
    ```

### artifact_tags Table Queries

#### InsertArtifactTag
- Description: Insert a new artifact tag into the "artifact_tags" table.
- SQL:
    ```sql
    INSERT INTO artifact_tags (
      artifact_id,
      tag_label,
      tag_value
    ) VALUES (
      $1,
      $2,
      $3
    ) RETURNING id;
    ```

#### GetArtifactTagByID
- Description: Retrieve an artifact tag from the "artifact_tags" table by ID.
- SQL:
    ```sql
    SELECT *
    FROM artifact_tags
    WHERE id = $1;
    ```

#### UpdateArtifactTag
- Description: Update an existing artifact tag in the "artifact_tags" table.
- SQL:
    ```sql
    UPDATE artifact_tags
    SET
      artifact_id = $1,
      tag_label = $2,
      tag_value = $3
    WHERE id = $4;
    ```

#### DeleteArtifactTag
- Description: Delete an artifact tag from the "artifact_tags" table.
- SQL:
    ```sql
    DELETE FROM artifact_tags
    WHERE id = $1;
    ```
