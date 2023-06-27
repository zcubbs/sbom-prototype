package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_CreateScanJobTx(t *testing.T) {
	store := NewStore(testDB)

	scanJob := createRandomScanJob(t)

	errs := make(chan error)
	results := make(chan CreateScanJobParamsTxResult)

	go func() {
		result, err := store.CreateScanJobTx(context.Background(), CreateScanJobParamsTx{
			InsertScanJobParams: InsertScanJobParams{
				ArtifactName:    scanJob.ArtifactName,
				ArtifactVersion: scanJob.ArtifactVersion,
				ArtifactType:    scanJob.ArtifactType,
				Status:          scanJob.Status,
			},
		})

		errs <- err
		results <- result
	}()

	err := <-errs

	require.NoError(t, err)

	result := <-results
	require.NotEmpty(t, result.ScanJob)

	resultScanJob := result.ScanJob
	require.NotEmpty(t, resultScanJob)
	require.Equal(t, scanJob.ArtifactName, resultScanJob.ArtifactName)
	require.Equal(t, scanJob.ArtifactVersion, resultScanJob.ArtifactVersion)
	require.Equal(t, scanJob.ArtifactType, resultScanJob.ArtifactType)
	require.Equal(t, scanJob.Status, resultScanJob.Status)
	require.NotZerof(t, resultScanJob.ID, "id should not be zero")
	require.NotEmpty(t, resultScanJob.CreatedAt, "created_at should not be empty")
	require.NotEmpty(t, resultScanJob.UpdatedAt, "updated_at should not be empty")

	_, err = store.GetScanJobByID(context.Background(), resultScanJob.ID)
	require.NoError(t, err)

	close(errs)
	close(results)
}
