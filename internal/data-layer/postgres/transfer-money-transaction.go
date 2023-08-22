package postgres

import (
	"context"
	"fmt"
	"gobanking/internal/common/types"
	"gobanking/internal/core-layer/domain"
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

	// call the ManageTransaction function to start and execute a money transfer transaction and pass it a callback function to be executed (this cb func is the actual transaction steps)
	err = ts.ManageTransaction(ctx, func(pg_tx *postgres.PG_TX) error {
		// => Initialize the repos concrete implementation via the transaction that is started via the transaction manager so all queries run within this same transaction
		account_repo := NewPG_AccountRepo(pg_tx)
		entry_repo := NewPG_EntryRepo(pg_tx)
		transfer_repo := NewPG_TransferRepo(pg_tx)

		// retrieve the account which will transfer the money to others to check if it's balance is sufficient for this operation
		fromAcc, err := account_repo.GetByID(args.FromAccountID)
		if err != nil {
			// return the error to the transaction manager to rollback the transaction
			return err
		}
		if fromAcc.Balance < args.Amount {
			// return the error to the transaction manager to rollback the transaction
			return fmt.Errorf("the current balance of this account with ownername : %v is smaller than the specified amount in the transactio", fromAcc.OwnerName)
		}

		// create transfer record for this money-transfer-tx
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
