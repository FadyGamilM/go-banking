package transaction

import (
	"context"
	"testing"

	account "github.com/FadyGamilM/go-banking-v2/internal/account/domain"

	"github.com/stretchr/testify/require"
)

func TestTransferMoneyTransaction(t *testing.T) {
	// define the context of the entire transaction (will be passed to the data layer via the service layer later when we build it)
	ctx := context.Background()

	// define the pgTxStore instance to manage our transaction
	pgTxStore := newPgTxStore(TestSqlDB)

	// create the account that will transfer the moeny
	FromAcc, err := createAccountForTest(ctx, &account.Account{
		OwnerName: "fady",
		Balance:   float64(200),
		Currency:  "EGP",
	})
	require.NoError(t, err)
	require.NotEmpty(t, FromAcc)
	require.NotZero(t, FromAcc.ID)

	// create the account that will receive the transfered money
	ToAcc, err := createAccountForTest(ctx, &account.Account{
		OwnerName: "samy",
		Balance:   float64(100),
		Currency:  "EGP",
	})
	require.NoError(t, err)
	require.NotEmpty(t, ToAcc)
	require.NotZero(t, ToAcc.ID)

	t.Logf(">> the initial value of balance before all transactions is $%v for the from-account and $%v for the to-account", FromAcc.Balance, ToAcc.Balance)

	// amount to be transfered
	amount := float64(10)

	// how many transaction will happen at almost the same time
	concurrentTXs := 5

	// channels so we can communicate with the running concurrent transactions
	errs := make(chan error)
	results := make(chan TransferMoneyTransactionResult)

	// run the concurrent transaction to heavily test my transaction implementation logic
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

	// test the results to ensure the correctnes of our transactions
	for i := 0; i < concurrentTXs; i++ {
		// read from the errors chanel to check if there is any error occurred
		err := <-errs
		require.NoError(t, err)

		// read the result from results channel
		result := <-results

		// test the transfer record
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, ToAcc.ID, transfer.ToAccountID)
		require.Equal(t, FromAcc.ID, transfer.FromAccountID)
		require.Equal(t, amount, transfer.Amount)

		// check the entry created for the to-account (which will record that there is a deposite has been occurred to this account)
		toAccEntry := result.ToEntry
		require.NotEmpty(t, toAccEntry)
		require.Equal(t, ToAcc.ID, toAccEntry.AccountID)

		// check the entry created for the from-account (which will record that there is a withdraw has been occurred from this account)
		fromAccEntry := result.FromEntry
		require.NotEmpty(t, fromAccEntry)
		require.Equal(t, FromAcc.ID, fromAccEntry.AccountID)

		// Start test-driven-development for the entire of the feature
		// -> we need to check the update that have beed done on both the toAcc and fromAcc
		updatedToAcc := result.ToAcc
		require.NotEmpty(t, updatedToAcc)
		require.NotZero(t, updatedToAcc.ID)
		require.NotZero(t, updatedToAcc.CreatedAt)
		// calculate the difference between the balance before and after the transaction
		// the toAcc instance i have created above "for testing" track the old balance before update and the updatedToAcc instance i just retrieved track the balance after update
		balanceDiff1 := ToAcc.Balance - updatedToAcc.Balance // this difference must be +ve or at least zero but it can't be -ve
		require.Positive(t, balanceDiff1)

		updatedFromAcc := result.FromAcc
		require.NotEmpty(t, updatedFromAcc)
		require.NotZero(t, updatedFromAcc.ID)
		require.NotZero(t, updatedFromAcc.CreatedAt)
		// calculate the difference between the balance before and after the transaction
		// the fromAcc instance tracks the old balance before the money-in, the updatedFromAcc instance tracks the updated balance after the money is transfered to the account
		balanceDiff2 := updatedFromAcc.Balance - FromAcc.Balance
		require.Positive(t, balanceDiff2)

		// both diferences must be the same
		require.Equal(t, balanceDiff1, balanceDiff2)

		t.Logf(">> the final value of balance after transaction #%v is $%v for the from-account and $%v for the to-account", i, updatedFromAcc.Balance, updatedToAcc.Balance)

	}

	// now we need to check the final balance of both accounts from the database
	finalUpdatedFromAcc, err := accTestRepo.GetByID(ctx, FromAcc.ID)
	require.NoError(t, err)
	require.Equal(t, FromAcc.Balance-(amount*float64(concurrentTXs)), finalUpdatedFromAcc.Balance)

	finalUpdatedToAcc, err := accTestRepo.GetByID(ctx, ToAcc.ID)
	require.NoError(t, err)
	require.Equal(t, ToAcc.Balance+(amount*float64(concurrentTXs)), finalUpdatedToAcc.Balance)

	t.Logf(">> the final value of balance after all transactions is $%v for the from-account and $%v for the to-account", finalUpdatedFromAcc.Balance, finalUpdatedToAcc.Balance)
}
