package transactions

import (
	"context"
	"database/sql"
	"fmt"
	"gobanking/internal/data-layer/store"
	"log"
)

/*
The transactionStore has 2 fields
 1. The Repos Store which consists of all repos of our application to access any CRUD operation of any resource
 2. The transaction type (TX) from sql driver to Begin | Rollback | Commit a transaction
*/
type transactionStore struct {
	Repos      *store.DataStore
	PG_DB_Conn *sql.DB
}

func NewTransactionStore(repos *store.DataStore, conn *sql.DB) *transactionStore {
	return &transactionStore{
		Repos:      repos,
		PG_DB_Conn: conn,
	}
}

func (ts *transactionStore) ManageTransaction(ctx context.Context, fn func(repoStore *store.DataStore) error) error {
	// begin the transaction
	tx, err := ts.PG_DB_Conn.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error while trying to BEGIN a transaction : %v \n ", err)
		return err
	}

	// now we will start executes queries within the context of this transaction (tx)
	err = fn(ts.Repos)
	// there is an error while executing one of the queries withint this transaction so we have to rollback
	if err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Printf("Error while trying to ROLLBACK the transaction : %v \n ", rollBackErr)
			return fmt.Errorf("Error occured in one of the queries within the transaction is : %v \n Error occurred during rolling back a transaction is : %v \n", err, rollBackErr)
		}
		return fmt.Errorf("Error occured in one of the queries within the transaction is : %v \n", err)
	}

	// if everything is okay  => Commit the transaction :D
	return tx.Commit() // tx.Commit() returns error or nil if there is no errors
}
