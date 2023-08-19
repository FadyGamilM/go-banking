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

func TestCreateEntry_ForBalanceIncrease(t *testing.T) {
	// teardown the database content to start fresh :D
	test_PG_DB.DB.Exec(`DELETE FROM entries`)
	test_PG_DB.DB.Exec(`DELETE FROM accounts`)

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
}

func TestCreateEntry_ForBalanceDecrease(t *testing.T) {
	// teardown the database content to start fresh :D
	test_PG_DB.DB.Exec(`DELETE FROM entries`)
	test_PG_DB.DB.Exec(`DELETE FROM accounts`)

	// first we create an account with initial balance
	acc_args := domain.Account{
		OwnerName: "Hossam",
		Balance:   float64(150.00),
		Currency:  "USD",
	}
	created_acc, err := createAccountForTest(&acc_args)
	require.NoError(t, err)

	// create a new entry in case user need to withdraw from his account (this will test the entry and update account both !)
	subtractionEntryBalance := float64(-350.00)
	entry_args := domain.Entry{
		AccountID: created_acc.ID,
		Amount:    subtractionEntryBalance,
	}
	created_entry, err := createEntryForTest(&entry_args)
	require.NoError(t, err)
	updated_account, err := test_acc_repo.UpdateByID(created_acc.ID, created_entry.Amount)
	require.Error(t, err)
	require.Nil(t, updated_account)

}

func TestGetAccountEntries(t *testing.T) {
	// teardown the database content to start fresh :D
	test_PG_DB.DB.Exec(`DELETE FROM entries`)
	test_PG_DB.DB.Exec(`DELETE FROM accounts`)

	acc_args := domain.Account{
		OwnerName: "Hossam",
		Balance:   float64(150.00),
		Currency:  "USD",
	}
	created_acc, err := createAccountForTest(&acc_args)
	require.NoError(t, err)

	additionEntryBalance := float64(+200.00)
	entry_args := domain.Entry{
		AccountID: created_acc.ID,
		Amount:    additionEntryBalance,
	}
	created_entry, err := createEntryForTest(&entry_args)
	require.NoError(t, err)

	// i assume the creation is tested successffully ..

	retrived_entry, err := test_entry_repo.GetOne(created_entry.ID, created_entry.AccountID)
	require.NoError(t, err)

	require.Equal(t, created_entry.ID, retrived_entry.ID)
	require.Equal(t, created_entry.Amount, retrived_entry.Amount)
	require.Equal(t, created_entry.AccountID, retrived_entry.AccountID)
	require.Equal(t, created_acc.ID, retrived_entry.AccountID)
}
