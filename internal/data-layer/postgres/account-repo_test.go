package postgres

import (
	"gobanking/internal/core-layer/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	req_args := domain.Account{
		OwnerName: "fady",
		Balance:   float64(100.00),
		Currency:  "USD",
	}

	created_acc, err := test_acc_repo.Create(&req_args)

	// now we need to check that

	// - No errors
	require.NoError(t, err)
	// - ID is not zero
	require.NotZero(t, created_acc.ID)
	// - Same provided fields are persisted in database
	require.Equal(t, req_args.OwnerName, created_acc.OwnerName)
	require.Equal(t, req_args.Balance, created_acc.Balance)
	require.Equal(t, req_args.Currency, created_acc.Currency)
	require.NotZero(t, created_acc.CreatedAt)
	require.NotZero(t, created_acc.UpdatedAt)
	require.Equal(t, true, created_acc.Activated)
}
