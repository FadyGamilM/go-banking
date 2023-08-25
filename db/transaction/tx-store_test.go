package transaction

import (
	"context"
	"testing"

	account "github.com/FadyGamilM/go-banking-v2/internal/account/domain"

	"github.com/stretchr/testify/require"
)

func TestTransferMoneyTransaction(t *testing.T) {
	ctx := context.Background()
	pgTxStore := newPgTxStore(TestSqlDB)

	FromAcc, err := createAccountForTest(ctx, &account.Account{
		OwnerName: "fady",
		Balance:   float64(200),
		Currency:  "EGP",
	})
	require.NoError(t, err)
	require.NotEmpty(t, FromAcc)
	require.NotZero(t, FromAcc.ID)

	ToAcc, err := createAccountForTest(ctx, &account.Account{
		OwnerName: "samy",
		Balance:   float64(100),
		Currency:  "EGP",
	})
	require.NoError(t, err)
	require.NotEmpty(t, ToAcc)
	require.NotZero(t, ToAcc.ID)

	t.Logf("the account (to) : %v", ToAcc.OwnerName)
	t.Logf("the account (from) : %v", FromAcc.OwnerName)

	amount := float64(10)
	concurrentTXs := 5
	errs := make(chan error)
	results := make(chan TransferMoneyTransactionResult)

	// result, err := pgTxStore.TransferMoneyTransaction(ctx, TransferMoneyTransactionParams{
	// 	ToAccID:   ToAcc.ID,
	// 	FromAccID: FromAcc.ID,
	// 	Amount:    amount,
	// })
	// require.NoError(t, err)

	// transfer := result.Transfer
	// require.NotEmpty(t, transfer)
	// require.Equal(t, ToAcc.ID, transfer.ToAccountID)
	// require.Equal(t, FromAcc.ID, transfer.FromAccountID)
	// require.Equal(t, amount, transfer.Amount)

	// toAccEntry := result.ToEntry
	// require.NotEmpty(t, toAccEntry)
	// require.Equal(t, ToAcc.ID, toAccEntry.AccountID)

	// fromAccEntry := result.FromEntry
	// require.NotEmpty(t, fromAccEntry)
	// require.Equal(t, FromAcc.ID, fromAccEntry.AccountID)

	for i := 0; i < concurrentTXs; i++ {
		go func() {
			result, err := pgTxStore.TransferMoneyTransaction(ctx, TransferMoneyTransactionParams{
				ToAccID:   ToAcc.ID,
				FromAccID: FromAcc.ID,
				Amount:    amount,
			})
			errs <- err
			results <- *result
		}()
	}

	for i := 0; i < concurrentTXs; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, ToAcc.ID, transfer.ToAccountID)
		require.Equal(t, FromAcc.ID, transfer.FromAccountID)
		require.Equal(t, amount, transfer.Amount)

		toAccEntry := result.ToEntry
		require.NotEmpty(t, toAccEntry)
		require.Equal(t, ToAcc.ID, toAccEntry.AccountID)

		fromAccEntry := result.FromEntry
		require.NotEmpty(t, fromAccEntry)
		require.Equal(t, FromAcc.ID, fromAccEntry.AccountID)

	}

}
