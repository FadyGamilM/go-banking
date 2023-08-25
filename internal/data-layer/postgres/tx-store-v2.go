package postgres

import (
	"context"
	"fmt"
	"gobanking/internal/common/types"
	"gobanking/internal/core-layer/domain"
	"gobanking/internal/infra-layer/db/postgres"
	"log"
)

type transactionsStoreV2 struct {
	// compose the PG_DB to extend the functionality of the sql.DB wrapper
	store_tx *postgres.PG_TX
	// contains the repos of the application to use their functionality
	Account_repo  *PG_AccountRepository
	Entry_repo    *PG_EntryRepository
	Transfer_repo *PG_TransferRepository
}

func NewTransactionsStoreV2(pgConn *postgres.PG_DB) *transactionsStoreV2 {

	tx, err := pgConn.DB.BeginTx(context.Background(), nil)
	if err != nil {
		log.Println("error while trying to begin the transaction and initiate a new transactions store => ", err)
	}

	pg_tx := &postgres.PG_TX{TX: tx}

	return &transactionsStoreV2{
		store_tx:      pg_tx,
		Account_repo:  NewPG_AccountRepo(pg_tx),
		Entry_repo:    NewPG_EntryRepo(pg_tx),
		Transfer_repo: NewPG_TransferRepo(pg_tx),
	}
}

func (tsv2 *transactionsStoreV2) TransferMoneyTX(tx_args *types.TransferMoneyTransactionParam) (*types.TransferMoneyTransactionResult, error) {
	// var txResult *types.TransferMoneyTransactionResult
	txResult := new(types.TransferMoneyTransactionResult)
	var err error

	ctx, cancel := CreateContext()
	defer cancel()

	// select the account to ensure that this account has an enough balance to make a transfer to another account
	// var fromAccBalance float64
	fromAccBalance := new(float64)
	log.Printf("the from account id is : %v \n", tx_args.FromAccountID)
	err = tsv2.store_tx.TX.QueryRowContext(ctx, `SELECT balance FROM accounts WHERE id = $1`, tx_args.FromAccountID).Scan(&fromAccBalance)
	log.Printf(" the retrieved balance is : %v\n", *fromAccBalance)
	if *fromAccBalance < tx_args.Amount {
		// rollback the transaction
		log.Printf("transaction cannot be completed because the balance is %v which is less than the specified amount => %v \n", fromAccBalance, err)
		rollBackErr := tsv2.store_tx.TX.Rollback()
		if rollBackErr != nil {
			log.Println("error while trying to rollback the transaction => ", rollBackErr)
			return nil, fmt.Errorf("error from the transaction => %v , error from the rollback => %v", err, rollBackErr)
		}
		return nil, err
	}

	// create a transafer record
	txResult.Transfer, err = tsv2.Transfer_repo.Create(&domain.Transfer{
		FromAccountID: tx_args.FromAccountID,
		ToAccountID:   tx_args.ToAccountID,
		Amount:        tx_args.Amount,
	})
	if err != nil {
		log.Println("error while trying to create a transfer record within the transaction => ", err)
		rollBackErr := tsv2.store_tx.TX.Rollback()
		if rollBackErr != nil {
			log.Println("error while trying to rollback the transaction => ", rollBackErr)
			return nil, fmt.Errorf("error from the transaction => %v , error from the rollback => %v", err, rollBackErr)
		}
		return nil, err
	}

	// create entry record for from-account
	txResult.FromEntry, err = tsv2.Entry_repo.Create(&domain.Entry{
		AccountID: tx_args.FromAccountID,
		Amount:    -tx_args.Amount,
	})
	if err != nil {
		log.Println("error while trying to create an entry record for the from-account within the transaction => ", err)
		rollBackErr := tsv2.store_tx.TX.Rollback()
		if rollBackErr != nil {
			log.Println("error while trying to rollback the transaction => ", rollBackErr)
			return nil, fmt.Errorf("error from the transaction => %v , error from the rollback => %v", err, rollBackErr)
		}
		return nil, err
	}

	// create entry record for to-account
	txResult.ToEntry, err = tsv2.Entry_repo.Create(&domain.Entry{
		AccountID: tx_args.ToAccountID,
		Amount:    tx_args.Amount,
	})
	if err != nil {
		log.Println("error while trying to create an entry record for the to-account within the transaction => ", err)
		rollBackErr := tsv2.store_tx.TX.Rollback()
		if rollBackErr != nil {
			log.Println("error while trying to rollback the transaction => ", rollBackErr)
			return nil, fmt.Errorf("error from the transaction => %v , error from the rollback => %v", err, rollBackErr)
		}
		return nil, err
	}

	// update the two accounts
	return txResult, nil
}
