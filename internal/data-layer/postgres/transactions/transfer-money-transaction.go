package transactions

import (
	"context"
	"gobanking/internal/common/types"
	"gobanking/internal/core-layer/domain"
	data_layer "gobanking/internal/data-layer/postgres"
	"gobanking/internal/infra-layer/db/postgres"
)

/*
➜ Fetch the account which will transfer the money to check it's balance

➜ Create new transfer record in database with the essential data info

➜ Create new entry record to track both money-in and money-out from both accounts

➜ Update the balance of both accounts
*/
func (ts *transactionStore) TransferMoneyTX(ctx context.Context, args types.TransferMoneyTransactionParam) (*types.TransferMoneyTransactionResult, error) {
	var txResult *types.TransferMoneyTransactionResult
	var err error

	err = ts.ManageTransaction(ctx, func(pg_tx *postgres.PG_TX) error {
		// account_repo := data_layer.NewPG_AccountRepo(pg_tx)
		entry_repo := data_layer.NewPG_EntryRepo(pg_tx)
		transfer_repo := data_layer.NewPG_TransferRepo(pg_tx)

		txResult.Transfer, err = transfer_repo.Create(&domain.Transfer{
			ToAccountID:   args.ToAccountID,
			FromAccountID: args.FromAccountID,
			Amount:        args.Amount,
		})
		if err != nil {
			// return the error to the transaction manager to rollback the transaction
			return err
		}

		txResult.FromEntry, err = entry_repo.Create(&domain.Entry{
			AccountID: args.FromAccountID,
			Amount:    -args.Amount,
		})
		if err != nil {
			// return the error to the transaction manager to rollback the transaction
			return err
		}

		txResult.ToEntry, err = entry_repo.Create(&domain.Entry{
			AccountID: args.ToAccountID,
			Amount:    args.Amount,
		})
		if err != nil {
			// return the error to the transaction manager to rollback the transaction
			return err
		}

		// TODO : Update both accounts later when we explore the deadlock solutions

		// if we here, so there is no error
		return nil
	})

	return txResult, err
}
