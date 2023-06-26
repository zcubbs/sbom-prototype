### scan_jobs Table Queries

#### InsertScanJob
- Description: Insert a new scan job into the "scan_jobs" table.
- SQL:
    ```sql
    INSERT INTO scan_jobs (
      artifact_uuid,
      artifact_name,
      artifact_version,
      artifact_type,
      status,
      report,
      job_log
    ) VALUES (
      $1,
      $2,
      $3,
      $4,
      $5,
      $6,
      $7
    ) RETURNING id;
    ```

#### GetScanJobByID
- Description: Retrieve a scan job from the "scan_jobs" table by ID.
- SQL:
    ```sql
    SELECT *
    FROM scan_jobs
    WHERE id = $1;
    ```

#### UpdateScanJob
- Description: Update an existing scan job in the "scan_jobs" table.
- SQL:
    ```sql
    UPDATE scan_jobs
    SET
      artifact_uuid = $1,
      artifact_name = $2,
      artifact_version = $3,
      artifact_type = $4,
      status = $5,
      report = $6,
      job_log = $7,
      updated_at = current_timestamp
    WHERE id = $8;
    ```

#### DeleteScanJob
- Description: Delete a scan job from the "scan_jobs" table.
- SQL:
    ```sql
    DELETE FROM scan_jobs
    WHERE id = $1;
    ```

#### GetScanJobsList
- Description: Retrieve a list of scan jobs from the "scan_jobs" table with pagination.
- SQL:
    ```sql
    SELECT *
    FROM scan_jobs
    ORDER BY id
    LIMIT $1
    OFFSET $2;
    ```
