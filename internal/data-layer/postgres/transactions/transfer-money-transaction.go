package transactions

import (
	"context"
	"log"
)

func (ts *transactionStore) TransferMoneyTX(ctx context.Context) error {
	// begin the transaction
	tx, err := ts.PG_DB_Conn.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error while trying to BEGIN a transaction : %v \n ", err)
		return err
	}

	// now we will start executes queries within the context of this transaction (tx)

}
