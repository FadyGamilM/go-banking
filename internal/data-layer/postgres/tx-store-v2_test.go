package postgres

import (
	"gobanking/internal/common/types"
	"gobanking/internal/core-layer/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferMoneyTX(t *testing.T) {

	toAcc, err := createAccountForTest(&domain.Account{
		OwnerName: "fady",
		Balance:   float64(200),
		Currency:  "USD",
	})
	require.NoError(t, err)
	require.NotEmpty(t, toAcc)
	require.NotZero(t, toAcc.ID)
	t.Logf("create the to-account with id : %v , owner name : %v , balance : %v \n", toAcc.ID, toAcc.OwnerName, toAcc.Balance)

	fromAcc, err := createAccountForTest(&domain.Account{
		OwnerName: "samy",
		Balance:   float64(200),
		Currency:  "USD",
	})
	require.NoError(t, err)
	require.NotEmpty(t, fromAcc)
	require.NotZero(t, fromAcc.ID)
	t.Logf("create the from-account with id : %v , owner name : %v , balance : %v \n", fromAcc.ID, fromAcc.OwnerName, fromAcc.Balance)

	amountToTransfer := float64(10)

	// time.Sleep(time.Duration(20 * time.Second))

	tx_args := &types.TransferMoneyTransactionParam{
		ToAccountID:   toAcc.ID,
		FromAccountID: fromAcc.ID,
		Amount:        amountToTransfer,
	}

	money_transfer_tx := NewTransactionsStoreV2(test_PG_DB)

	money_transfer_tx_result, err := money_transfer_tx.TransferMoneyTX(tx_args)
	require.Error(t, err)
	require.NotEmpty(t, money_transfer_tx_result)
	t.Logf("result of transaction.Transfer : %v", money_transfer_tx_result.Transfer.ID)
}
