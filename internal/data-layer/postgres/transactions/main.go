package transactions

import (
	"context"
	"fmt"
	account_ports "gobanking/internal/core-layer/ports/accounts"
	entry_ports "gobanking/internal/core-layer/ports/entries"
	transfer_ports "gobanking/internal/core-layer/ports/transfers"
	"gobanking/internal/data-layer/store"
	"gobanking/internal/infra-layer/db/postgres"
)

// TxStore is the store that is used to execute entire transaction
type TxStore struct {
	PG         *postgres.PG_DB
	Repo_Store *store.DataStore
}

func NewTxStore(db *postgres.PG_DB, acc *account_ports.AccountRepository, entry *entry_ports.EntryRepository, transfer *transfer_ports.TransferRepository) *TxStore {
	return &TxStore{
		PG:         db,
		Repo_Store: store.NewDataStore(acc, entry, transfer),
	}
}

func (txStore *TxStore) execTx(ctx context.Context, fn func(pg_db *postgres.PG_DB) error) error {
	// begin the transaction
	tx, err := txStore.PG.DB.BeginTx(ctx, nil)
	// if there is an error while beginning the transaction, thats means there is no transaction is started, so we return and end
	if err != nil {
		return err
	}

	// execute the db function (transaction or part of the transaction)
	err = fn(txStore.PG)
	if err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			return fmt.Errorf("transaction error => %v , rollback error => %v", err, rollBackErr)
		}
		// if rollback is done successfully, return the transaction error
		return err
	}

	// if all operatons are done => commit
	return tx.Commit()
}

// transfer money transaction ..
func (txStore *TxStore) TransferMoneyTransaction() {}
