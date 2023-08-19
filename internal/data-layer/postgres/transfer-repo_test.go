package postgres

import (
	"gobanking/internal/core-layer/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func createAccount(acc_args *domain.Account) (*domain.Account, error) {
	return test_acc_repo.Create(acc_args)
}

func TestCreateTransfer(t *testing.T) {
	// 0. teardown the database tables
	test_PG_DB.DB.Exec(`DELETE FROM transfers`)
	test_PG_DB.DB.Exec(`DELETE FROM accounts`)

	// 1. create two accounts
	acc_1_args := domain.Account{OwnerName: "fady gamil", Balance: float64(190.66), Currency: "EUR"}
	acc_1, err := createAccount(&acc_1_args)
	require.NoError(t, err) // we already tested the account-resource crud operations so i won't validate anything here related to the account-resource

	acc_2_args := domain.Account{OwnerName: "nader nabil", Balance: float64(200.00), Currency: "EUR"}
	acc_2, err := createAccount(&acc_2_args)
	require.NoError(t, err) // we already tested the account-resource crud operations so i won't validate anything here related to the account-resource

	// 2. create a transfer instance
	amount_to_transfer := float64(100.00)
	transfer_args := domain.Transfer{ToAccountID: acc_1.ID, FromAccountID: acc_2.ID, Amount: amount_to_transfer}
	created_transfer, err := test_transfer_Repo.Create(&transfer_args)
	require.NoError(t, err)
	require.Equal(t, created_transfer.Amount, amount_to_transfer)
	require.Equal(t, created_transfer.ToAccountID, acc_1.ID)
	require.Equal(t, created_transfer.FromAccountID, acc_2.ID)
	require.NotZero(t, created_transfer.CreatedAt)
	require.NotZero(t, created_transfer.UpdatedAt)
}
