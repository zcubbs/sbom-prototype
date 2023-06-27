package db

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tabbed/pqtype"
	"testing"
	"zel/sbom-prototype/scanner/util"
)

func createRandomScanJob(t *testing.T) ScanJob {
	arg := InsertScanJobParams{
		ArtifactUuid:    uuid.NullUUID{},
		ArtifactName:    util.RandomArtifactName(),
		ArtifactVersion: util.RandomArtifactVersion(),
		ArtifactType:    util.RandomArtifactType(),
		Status:          "created",
		Report:          pqtype.NullRawMessage{},
		JobLog:          sql.NullString{},
	}
	job, err := testQueries.InsertScanJob(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, job)

	require.NotEmpty(t, arg.ArtifactName, "artifact name should not be empty")
	require.NotEmpty(t, arg.ArtifactVersion, "artifact version should not be empty")
	require.NotEmpty(t, arg.ArtifactType, "artifact type should not be empty")
	require.Equal(t, job.Status, "created")
	require.NotZerof(t, job.ID, "id should not be zero")
	require.NotEmpty(t, job.CreatedAt, "created_at should not be empty")
	require.NotEmpty(t, job.UpdatedAt, "updated_at should not be empty")

	return job
}

func TestQueries_InsertScanJob(t *testing.T) {
	scanJob := createRandomScanJob(t)

	require.NotEmpty(t, scanJob)

	require.NotEmpty(t, scanJob.ArtifactName, "artifact name should not be empty")
	require.NotEmpty(t, scanJob.ArtifactVersion, "artifact version should not be empty")
	require.NotEmpty(t, scanJob.ArtifactType, "artifact type should not be empty")
	require.Equal(t, scanJob.Status, "created")
	require.NotZerof(t, scanJob.ID, "id should not be zero")
	require.NotEmpty(t, scanJob.CreatedAt, "created_at should not be empty")
	require.NotEmpty(t, scanJob.UpdatedAt, "updated_at should not be empty")
}

func TestQueries_GetScanJob(t *testing.T) {
	scanJob := createRandomScanJob(t)

	got, err := testQueries.GetScanJobByID(context.Background(), scanJob.ID)
	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.NotEmpty(t, got.ArtifactName, "artifact name should not be empty")
	require.NotEmpty(t, got.ArtifactVersion, "artifact version should not be empty")
	require.NotEmpty(t, got.ArtifactType, "artifact type should not be empty")
	require.Equal(t, got.Status, "created")
	require.NotZerof(t, got.ID, "id should not be zero")
	require.NotEmpty(t, got.CreatedAt, "created_at should not be empty")
	require.NotEmpty(t, got.UpdatedAt, "updated_at should not be empty")
}

func TestQueries_UpdateScanJob(t *testing.T) {
	scanJob := createRandomScanJob(t)

	arg := UpdateScanJobParams{
		ArtifactUuid:    uuid.NullUUID{},
		ArtifactName:    util.RandomArtifactName(),
		ArtifactVersion: util.RandomArtifactVersion(),
		ArtifactType:    util.RandomArtifactType(),
		Status:          "scanning",
		Report:          pqtype.NullRawMessage{},
		JobLog:          sql.NullString{},
		ID:              scanJob.ID,
	}

	got, err := testQueries.UpdateScanJob(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.NotEmpty(t, got.ArtifactName, "artifact name should not be empty")
	require.NotEmpty(t, got.ArtifactVersion, "artifact version should not be empty")
	require.NotEmpty(t, got.ArtifactType, "artifact type should not be empty")
	require.Equal(t, got.Status, "scanning")
	require.NotZerof(t, got.ID, "id should not be zero")
	require.NotEmpty(t, got.CreatedAt, "created_at should not be empty")
	require.NotEmpty(t, got.UpdatedAt, "updated_at should not be empty")
}

func TestQueries_DeleteScanJob(t *testing.T) {
	scanJob := createRandomScanJob(t)

	err := testQueries.DeleteScanJob(context.Background(), scanJob.ID)
	require.NoError(t, err)

	got, err := testQueries.GetScanJobByID(context.Background(), scanJob.ID)
	require.Error(t, err)
	require.Empty(t, got)
}

func TestQueries_ListScanJobs(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomScanJob(t)
	}

	arg := GetScanJobsListParams{
		Limit:  5,
		Offset: 5,
	}

	jobs, err := testQueries.GetScanJobsList(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, jobs, 5)

	for _, job := range jobs {
		require.NotEmpty(t, job.ArtifactName, "artifact name should not be empty")
		require.NotEmpty(t, job.ArtifactVersion, "artifact version should not be empty")
		require.NotEmpty(t, job.ArtifactType, "artifact type should not be empty")
		require.NotZerof(t, job.ID, "id should not be zero")
		require.NotEmpty(t, job.CreatedAt, "created_at should not be empty")
		require.NotEmpty(t, job.UpdatedAt, "updated_at should not be empty")
	}
}
