package postgres

import (
	"context"
	"gobanking/internal/common/types"
	"gobanking/internal/core-layer/domain"
	"gobanking/internal/data-layer/postgres/transactions"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferMoneyTransaction(t *testing.T) {
	// inistantiate a new instance of the transaction store type to execute the transaction
	txStore := transactions.NewTransactionStore(test_PG_DB)

	// define number of concurrent transactions that will run at the same time (approx)
	numConcurrentTXs := 5

	// define the amount to be transfered for each transaction
	amountToTransfer := float64(10)

	// define error channel and result channel so each go routine can communicate with the main routine (which is this test function) so we will be able to perform some require checks on the results
	errsChan := make(chan error)
	resultChan := make(chan *types.TransferMoneyTransactionResult)

	// define the from account and to account which will be the 2 parties of the transaction
	fromAcc, err := createAccountForTest(&domain.Account{
		OwnerName: "hossam abdu",
		Balance:   float64(100),
		Currency:  "EGB",
	})
	if err != nil {
		t.Logf("error while creating the from account : %v \n", err)
	}
	toAcc, err := createAccountForTest(&domain.Account{
		OwnerName: "fady gamil",
		Balance:   float64(100),
		Currency:  "EGB",
	})
	if err != nil {
		t.Logf("error while creating the to account : %v \n", err)
	}

	for tx_idx := 0; tx_idx < numConcurrentTXs; tx_idx++ {
		go func() {
			txResult, err := txStore.TransferMoneyTX(context.Background(), types.TransferMoneyTransactionParam{
				ToAccountID:   toAcc.ID,
				FromAccountID: fromAcc.ID,
				Amount:        amountToTransfer,
			})

			// now communicate to main routine and send the error and the result
			errsChan <- err
			resultChan <- txResult
		}()
	}

	// get the result to assert the test cases
	for res_idx := 0; res_idx < numConcurrentTXs; res_idx++ {
		// check if there is any error
		err := <-errsChan
		require.NoError(t, err)

		// get the tx result
		res := <-resultChan
		// no empty results
		require.NotEmpty(t, res.FromAccount)
		require.NotEmpty(t, res.FromEntry)
		require.NotEmpty(t, res.ToAccount)
		require.NotEmpty(t, res.ToEntry)
		require.NotEmpty(t, res.Transfer)

		created_transafer := res.Transfer
		require.Equal(t, created_transafer.ID, )
	}
}
