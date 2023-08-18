package postgres

import (
	"gobanking/internal/core-layer/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func createEntryForTest(req_args *domain.Entry) (*domain.Entry, error) {
	return test_entry_repo.Create(req_args)
}

func createUserForTest(req_args *domain.Account) (*domain.Account, error) {

	created_acc, err := test_acc_repo.Create(req_args)

	return created_acc, err
}

func TestCreateEntry(t *testing.T) {
	// teardown the database content to start fresh :D
	test_PG_DB.DB.Exec(`DELETE FROM accounts`)
	test_PG_DB.DB.Exec(`DELETE FROM entries`)

	// first we create an account with initial balance
	acc_args := domain.Account{
		OwnerName: "mina",
		Balance:   float64(150.00),
		Currency:  "USD",
	}
	created_acc, err := createAccountForTest(&acc_args)
	require.NoError(t, err)

	// create a new entry in case user need to deposite to his account (this will test the entry and update account both !)
	additionEntryBalance := float64(+200.00)
	entry_args := domain.Entry{
		AccountID: created_acc.ID,
		Amount:    additionEntryBalance,
	}
	created_entry, err := createEntryForTest(&entry_args)
	require.NoError(t, err)
	updated_acc, err := test_acc_repo.UpdateByID(created_acc.ID, created_entry.Amount)
	require.NoError(t, err)
	require.Equal(t, updated_acc.Balance, created_acc.Balance+created_entry.Amount)
	require.NotZero(t, created_entry.ID)
	require.NotZero(t, created_entry.CreatedAt)
	require.NotZero(t, created_entry.UpdatedAt)

	// create a new entry in case user need to withdraw from his account (this will test the entry and update account both !)
	subtractionEntryBalance := float64(-350.00)
	entry_args = domain.Entry{
		AccountID: created_acc.ID,
		Amount:    subtractionEntryBalance,
	}
	created_entry, err = createEntryForTest(&entry_args)
	require.NoError(t, err)
	updated_acc, err = test_acc_repo.UpdateByID(created_acc.ID, created_entry.Amount)
	require.Error(t, err)
	// require.Equal(t, updated_acc.Balance, created_acc.Balance+created_entry.Amount) // + because we the amoutn is -ve
	// require.NotZero(t, created_entry.ID)
	// require.NotZero(t, created_entry.CreatedAt)
	// require.NotZero(t, created_entry.UpdatedAt)

}

func TestGetAccountEntries(t *testing.T) {

}
