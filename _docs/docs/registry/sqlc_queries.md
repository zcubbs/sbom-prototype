### sboms Table Queries

#### InsertScanJob
- Description: Insert a new scan job into the "sboms" table.
- SQL:
    ```sql
    INSERT INTO sboms (
      uuid,
      created_at,
      manifest,
      job_log,
      status,
      artifact_uuid,
      artifact_name,
      artifact_version,
      artifact_type
    ) VALUES (
      $1,
      $2,
      $3,
      $4,
      $5,
      $6,
      $7,
      $8,
      $9
    ) RETURNING id;
    ```

#### GetScanJobByID
- Description: Retrieve a scan job from the "sboms" table by ID.
- SQL:
    ```sql
    SELECT *
    FROM sboms
    WHERE id = $1;
    ```

#### UpdateScanJob
- Description: Update an existing scan job in the "sboms" table.
- SQL:
    ```sql
    UPDATE sboms
    SET
      uuid = $1,
      created_at = $2,
      manifest = $3,
      job_log = $4,
      status = $5,
      artifact_uuid = $6,
      artifact_name = $7,
      artifact_version = $8,
      artifact_type = $9,
      updated_at = current_timestamp
    WHERE id = $10;
    ```

#### DeleteScanJob
- Description: Delete a scan job from the "sboms" table.
- SQL:
    ```sql
    DELETE FROM sboms
    WHERE id = $1;
    ```

#### GetScanJobsList
- Description: Retrieve a list of scan jobs from the "sboms" table with pagination.
- SQL:
    ```sql
    SELECT *
    FROM sboms
    ORDER BY id
    LIMIT $1
    OFFSET $2;
    ```
