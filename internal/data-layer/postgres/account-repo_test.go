package postgres

import (
	"gobanking/internal/core-layer/domain"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createAccountForTest(req_args *domain.Account) (*domain.Account, error) {

	created_acc, err := test_acc_repo.Create(req_args)

	return created_acc, err
}

func TestCreateAccount(t *testing.T) {

	req_args := domain.Account{
		OwnerName: "fady",
		Balance:   float64(100.00),
		Currency:  "USD",
	}

	created_acc, err := createAccountForTest(&req_args)

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

func TestGetAccountByID(t *testing.T) {
	req_args := domain.Account{
		OwnerName: "ahmed",
		Balance:   float64(100.00),
		Currency:  "USD",
	}

	created_acc, err := createAccountForTest(&req_args)

	require.NoError(t, err)

	retrieved_acc, err := test_acc_repo.GetByID(created_acc.ID)
	log.Println(created_acc)
	log.Println(retrieved_acc)

	require.NoError(t, err)

	require.Equal(t, created_acc.ID, retrieved_acc.ID)
	require.Equal(t, created_acc.OwnerName, retrieved_acc.OwnerName)
	require.Equal(t, created_acc.Balance, retrieved_acc.Balance)
	require.Equal(t, created_acc.Currency, retrieved_acc.Currency)
	require.Equal(t, created_acc.Activated, retrieved_acc.Activated)
	require.WithinDuration(t, created_acc.CreatedAt, retrieved_acc.CreatedAt, time.Second)

	require.NotZero(t, retrieved_acc.UpdatedAt)
}

func TestGetAllAccounts(t *testing.T) {

	test_acc_repo.pg.DB.Exec(`DELETE FROM accounts`)

	req_args_1 := domain.Account{
		OwnerName: "ahmed",
		Balance:   float64(100.00),
		Currency:  "USD",
	}

	req_args_2 := domain.Account{
		OwnerName: "ahmed",
		Balance:   float64(100.00),
		Currency:  "USD",
	}

	_, err := createAccountForTest(&req_args_1)
	require.NoError(t, err)

	_, err = createAccountForTest(&req_args_2)

	require.NoError(t, err)

	retrieved_accounts, err := test_acc_repo.GetAll()
	log.Println(retrieved_accounts)

	require.NoError(t, err)

	t.Logf("all accounts are %v \n", len(retrieved_accounts))
	require.Len(t, retrieved_accounts, 2)
}

func TestDeleteAccountByID(t *testing.T) {
	test_acc_repo.pg.DB.Exec(`DELETE FROM accounts`)

	req_args := domain.Account{
		OwnerName: "ahmed",
		Balance:   float64(100.00),
		Currency:  "USD",
	}

	created_acc, err := createAccountForTest(&req_args)

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

	// delete the account
	err = test_acc_repo.DeleteByID(created_acc.ID)
	require.NoError(t, err)

	// get all accounts to ensure that there is no accounts in database with this id
	retrieved_accounts, err := test_acc_repo.GetAll()
	require.NoError(t, err)
	require.Len(t, retrieved_accounts, 0)
}

func TestUpdateAccountByID(t *testing.T) {
	test_acc_repo.pg.DB.Exec(`DELETE FROM accounts`)

	req_args := domain.Account{
		OwnerName: "ahmed",
		Balance:   float64(100.00),
		Currency:  "USD",
	}

	created_acc, err := createAccountForTest(&req_args)

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

	balanceUpdate := float64(200.0)
	updated_acc, err := test_acc_repo.UpdateByID(created_acc.ID, created_acc.Balance+balanceUpdate)

	// test that the balance is not the same and the right value is updated
	require.NoError(t, err)            // no errors
	require.NotZero(t, created_acc.ID) // id still the same and record still in db
	require.NotEqual(t, created_acc.Balance, updated_acc.Balance)
	require.Equal(t, updated_acc.Balance, created_acc.Balance+balanceUpdate)
}
