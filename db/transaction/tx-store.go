package transaction

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	database "github.com/FadyGamilM/go-banking-v2/db"
	account "github.com/FadyGamilM/go-banking-v2/internal/account/domain"
	entry "github.com/FadyGamilM/go-banking-v2/internal/entry/domain"
	transfer "github.com/FadyGamilM/go-banking-v2/internal/transfer/domain"

	account_repo "github.com/FadyGamilM/go-banking-v2/internal/account"
	entry_repo "github.com/FadyGamilM/go-banking-v2/internal/entry"
	transfer_repo "github.com/FadyGamilM/go-banking-v2/internal/transfer"
)

// this type provides for us the ability to run individual queries or combination of queries within a transaction
type PgTxStore struct {
	// lets now extend the PG type functionality by composing it
	*database.PG
	// to create a new db transaction
	db *sql.DB
}

// create new pg_tx_store object
func newPgTxStore(db *sql.DB) *PgTxStore {
	return &PgTxStore{
		db: db,
		PG: database.NewPG(db),
	}
}

var txKey = struct{}{}

// generate and execute  a database transaction
/*
	@ recieves :
		- context
		- callback function
	@ logic :
		- start a new database transaction (*sql.Tx)
		- create new PG object within the current transaction
		- call the callback function with the created PG instance
		- Commit the transaction
*/
func (pgtx *PgTxStore) execTransaction(ctx context.Context, txFn func(*database.PG) error) error {
	// TODO : set the isolation level as nil for now (modified later)
	tx, err := pgtx.db.BeginTx(ctx, nil)
	if err != nil {

		return fmt.Errorf("error while trying to begin the transaction : %v", err)
	}

	// now this newPgtx provides for us to run a database query within the transaction
	newPgtx := database.NewPG(tx)

	err = txFn(newPgtx)

	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			return fmt.Errorf("error while trying to begin the transaction : %v , error while trying to rollback the transation : %v ", err, rollBackErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferMoneyTransactionParams struct {
	ToAccID   int64   `json:"to_account_id"`
	FromAccID int64   `json:"from_account_id"`
	Amount    float64 `json:"amount"`
}

type TransferMoneyTransactionResult struct {
	ToAcc     *account.Account
	FromAcc   *account.Account
	ToEntry   *entry.Entry
	FromEntry *entry.Entry
	Transfer  *transfer.Transfer
}

func (pgtx *PgTxStore) TransferMoneyTransaction(ctx context.Context, args TransferMoneyTransactionParams) (*TransferMoneyTransactionResult, error) {
	result := new(TransferMoneyTransactionResult)

	err := pgtx.execTransaction(ctx, func(pgdb *database.PG) error {
		var err error

		// inside this callback function we are inside the tx transaction
		// and any running query via the `q` instance is running within the transaction
		// so i will instantiate the repos here to perform the transaction with this instance of pg conn
		accRepo := account_repo.NewAccountRepo(pgdb)
		entryRepo := entry_repo.NewEntryRepo(pgdb)
		transferRepo := transfer_repo.NewTransferRepo(pgdb)

		// extract the current transaction name from the passed context (for testing purpose only and from the context which is passed from the unit test only)
		txName := ctx.Value(txKey)

		// fetch the from-account to ensure that its stored in database and its balance is sufficient to perform this transafer
		// log.Printf("[%v] | fetch the from-account", txName)
		// retrievedFromAcc, err := accRepo.GetByID(ctx, args.FromAccID)
		// if err != nil {
		// 	log.Printf("error while fetching the from-account from database : %v", err)
		// 	return err
		// }
		// if retrievedFromAcc.ID == int64(0) {
		// 	return fmt.Errorf("there is no account with id = %v", retrievedFromAcc.ID)
		// }
		// if retrievedFromAcc.Balance < args.Amount {
		// 	return fmt.Errorf("the account with id %v has balance = $%v , which is less than the specified amount = %v in the transaction", retrievedFromAcc.ID, retrievedFromAcc.Balance, args.Amount)
		// }

		// // fetch the to-account to ensure that its stored in database
		// log.Printf("[%v] | fetch the to-account", txName)
		// retrievedToAcc, err := accRepo.GetByID(ctx, args.ToAccID)
		// if err != nil {
		// 	log.Printf("error while fetching the to-account from database : %v ", err)
		// 	return err
		// }
		// if retrievedToAcc.ID == int64(0) {
		// 	return fmt.Errorf("there is no account with id = %v", args.ToAccID)
		// }

		// create the transfer record in database
		log.Printf("[%v] | create the transfer record", txName)
		result.Transfer, err = transferRepo.Create(ctx, &transfer.Transfer{
			ToAccountID:   args.ToAccID,
			FromAccountID: args.FromAccID,
			Amount:        args.Amount,
		})
		if err != nil {
			return fmt.Errorf("error while creating a new transfer record : %v", err)
		}

		// create a to-entry record which represents the money-transfer-in (Additional money to the to-account)
		log.Printf("[%v] | create the to-entry record", txName)
		result.ToEntry, err = entryRepo.Create(ctx, &entry.Entry{
			AccountID: args.ToAccID,
			Amount:    +args.Amount,
		})
		if err != nil {
			return fmt.Errorf("error while creating a new to-entry record : %v", err)
		}

		// create a from-entry record which represents the money-transfer-out (reduce money from the from-account)
		log.Printf("[%v] | create the from-entry record", txName)
		result.FromEntry, err = entryRepo.Create(ctx, &entry.Entry{
			AccountID: args.FromAccID,
			Amount:    -args.Amount,
		})
		if err != nil {
			return fmt.Errorf("error while creating a new from-entry record : %v", err)
		}

		// TODO => update the accounts balances
		// reorder the update of the 2 accounts so we ensure that the change to acc1 in tx1 will be done and then the change to acc2 to tx1 is done after it, and at the same time the change in acc1 in tx2 is done before the change of the acc2 of tx2 is done even if in tx2 the change to acc2 is acquired first i will re-order it
		if args.FromAccID > args.ToAccID {
			// i already have fetched the accounts so i don't have to ensure that there are accounts with these provided ids .. now lets move to the update query
			log.Printf("[%v] | update the from-account", txName)
			updatedFromAcc, err := accRepo.UpdateByID(ctx, args.FromAccID, -args.Amount)
			if err != nil {
				return fmt.Errorf("error while updating the from-account balance : %v", err)
			}

			log.Printf("[%v] | update the to-account", txName)
			updatedToAcc, err := accRepo.UpdateByID(ctx, args.ToAccID, args.Amount)
			if err != nil {
				return fmt.Errorf("error while updating the to-account balance : %v", err)
			}
			result.FromAcc = updatedFromAcc
			result.ToAcc = updatedToAcc
		} else {
			log.Printf("[%v] | update the to-account", txName)
			updatedToAcc, err := accRepo.UpdateByID(ctx, args.ToAccID, args.Amount)
			if err != nil {
				return fmt.Errorf("error while updating the to-account balance : %v", err)
			}

			log.Printf("[%v] | update the from-account", txName)
			updatedFromAcc, err := accRepo.UpdateByID(ctx, args.FromAccID, -args.Amount)
			if err != nil {
				return fmt.Errorf("error while updating the from-account balance : %v", err)
			}
			result.FromAcc = updatedFromAcc
			result.ToAcc = updatedToAcc
		}

		return nil
	})

	return result, err
}
