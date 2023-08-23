package postgres

import (
	"context"
	"gobanking/internal/common/types"
	"testing"
)

func TestTransferMoneyTransaction(t *testing.T) {
	// // 0. teardown the database tables
	// test_PG_DB.DB.Exec(`DELETE FROM transfers`)
	// test_PG_DB.DB.Exec(`DELETE FROM accounts`)

	// inistantiate a new instance of the transaction store type to execute the transaction
	txStore := NewTransactionStore(test_PG_DB)

	// define number of concurrent transactions that will run at the same time (approx)
	// numConcurrentTXs := 5

	// define the amount to be transfered for each transaction
	amountToTransfer := float64(10)

	// var fromAcc, toAcc *domain.Account
	// var err error

	// define error channel and result channel so each go routine can communicate with the main routine (which is this test function) so we will be able to perform some require checks on the results
	// errsChan := make(chan error)
	// resultChan := make(chan *types.TransferMoneyTransactionResult)

	// // define the from account and to account which will be the 2 parties of the transaction
	// fromAcc, err = createAccountForTest(&domain.Account{
	// 	OwnerName: "hossam abdu",
	// 	Balance:   float64(100),
	// 	Currency:  "EGB",
	// })
	// if err != nil {
	// 	t.Logf("error while creating the from account : %v \n", err)
	// }

	// toAcc, err = createAccountForTest(&domain.Account{
	// 	OwnerName: "fady gamil",
	// 	Balance:   float64(100),
	// 	Currency:  "EGB",
	// })
	// if err != nil {
	// 	t.Logf("error while creating the to account : %v \n", err)
	// }

	// time.Sleep(time.Duration(time.Duration.Seconds(10)))

	_, err := txStore.TransferMoneyTX(context.Background(), &types.TransferMoneyTransactionParam{
		ToAccountID:   int64(51),
		FromAccountID: int64(52),
		Amount:        amountToTransfer,
	})

	t.Log("the error due to transaction is : => ", err)
	// t.Log("the transaction result is : => ", txResult.FromAccount, "\n", txResult.ToAccount, "\n", txResult.Transfer)

	// for tx_idx := 0; tx_idx < numConcurrentTXs; tx_idx++ {
	// 	go func() {
	// 		log.Println("inside go routine, the id of the from account is : ", int64(52))
	// 		txResult, err := txStore.TransferMoneyTX(context.Background(), &types.TransferMoneyTransactionParam{
	// 			ToAccountID:   int64(51),
	// 			FromAccountID: int64(52),
	// 			Amount:        amountToTransfer,
	// 		})

	// 		// now communicate to main routine and send the error and the result
	// 		errsChan <- err
	// 		resultChan <- txResult
	// 	}()
	// }

	// // get the result to assert the test cases
	// for res_idx := 0; res_idx < numConcurrentTXs; res_idx++ {
	// 	// check if there is any error
	// 	err := <-errsChan
	// 	require.NoError(t, err)

	// 	// get the tx result
	// 	res := <-resultChan
	// 	// no empty results
	// 	require.NotEmpty(t, res.FromAccount)
	// 	require.NotEmpty(t, res.FromEntry)
	// 	require.NotEmpty(t, res.ToAccount)
	// 	require.NotEmpty(t, res.ToEntry)
	// 	require.NotEmpty(t, res.Transfer)

	// 	// check the created transfer record against our expected values
	// 	created_transafer := res.Transfer
	// 	require.NotZero(t, created_transafer.ID)
	// 	require.Equal(t, int64(52), created_transafer.FromAccountID)
	// 	require.Equal(t, int64(51), created_transafer.ToAccountID)
	// 	require.Equal(t, amountToTransfer, created_transafer.Amount)
	// 	require.NotZero(t, created_transafer.CreatedAt)
	// 	require.NotZero(t, created_transafer.UpdatedAt)
	// 	// and them retrieve the created transfer record from the database to ensure that its persisted and the data are true
	// 	_, err = test_transfer_Repo.GetByID(created_transafer.ID)
	// 	require.NoError(t, err)

	// 	// check the created entries
	// 	from_acc_entry := res.FromEntry
	// 	require.NotZero(t, from_acc_entry.ID)
	// 	require.Equal(t, int64(52), from_acc_entry.AccountID)
	// 	// the entry is created using the given amount as +ve and at the time of creatinon we pass the -ve sign in case of the from-account, so we have to expect the -ve value with the amountToTransfer too
	// 	require.Equal(t, -amountToTransfer, from_acc_entry.Amount)
	// 	require.NotZero(t, from_acc_entry.CreatedAt)
	// 	require.NotZero(t, from_acc_entry.UpdatedAt)

	// 	to_acc_entry := res.ToEntry
	// 	require.NotZero(t, to_acc_entry.ID)
	// 	require.Equal(t, int64(51), to_acc_entry.AccountID)
	// 	require.Equal(t, amountToTransfer, to_acc_entry.Amount)
	// 	require.NotZero(t, to_acc_entry.CreatedAt)
	// 	require.NotZero(t, to_acc_entry.UpdatedAt)

	// TODO : test the updated accounts when we utilize the concurrency patterns

	// }
}
