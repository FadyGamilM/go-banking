package postgres

import (
	"context"
	"fmt"
	"gobanking/internal/infra-layer/db/postgres"
	"log"
)

/*
The transactionStore has 2 fields
 1. The Repos Store which consists of all repos of our application to access any CRUD operation of any resource
 2. The transaction type (TX) from sql driver to Begin | Rollback | Commit a transaction

This type provides for us the ability to run individual queries and combined queries to form complete transaction
*/
type transactionStore struct {
	PG_Conn *postgres.PG_DB
}

func NewTransactionStore(conn_pool *postgres.PG_DB) *transactionStore {
	// create new instance of the transaction store type
	return &transactionStore{PG_Conn: conn_pool}
}

func (ts *transactionStore) ManageTransaction(ctx context.Context, fn func(pg_conn *postgres.PG_TX) error) error {
	pg_tx, err := ts.PG_Conn.DB.BeginTx(ctx, nil) // nil for now but later i will specify the required isolation context
	if err != nil {
		log.Printf("Error while trying to BEGIN a transaction : %v \n ", err)
		return err
	}

	// now we will start executes queries within the context of this transaction (tx)
	// so we need to create new instance of *sql.DB and pass it to the concrete implementation of all our repos implementations so all of them will execute their built CRUD methods using same instance within the current transaction isolated from other transactions instances
	tx := *&postgres.PG_TX{TX: pg_tx}
	err = fn(&tx)
	// there is an error while executing one of the queries withint this transaction so we have to rollback
	if err != nil {
		if rollBackErr := tx.TX.Rollback(); rollBackErr != nil {
			log.Printf("Error while trying to ROLLBACK the transaction : %v \n ", rollBackErr)
			return fmt.Errorf("Error occured in one of the queries within the transaction is : %v \n Error occurred during rolling back a transaction is : %v \n", err, rollBackErr)
		}
		return fmt.Errorf("Error occured in one of the queries within the transaction is : %v \n", err)
	}

	// if everything is okay  => Commit the transaction :D
	return tx.TX.Commit() // tx.Commit() returns error or nil if there is no errors
}
