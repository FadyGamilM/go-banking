package postgres

// import (
// 	"gobanking/internal/core/domain"
// 	"log"
// )

// type TransactionStore struct {
// 	// composition to extend the functionality of the postgres_repo
// 	*Postgres_repo // because i might need to execute custom sql queries in the middle of transaction which is not implemented in the next added repos
// 	acc_repo       Account_repo
// 	trans_repo     Transfer_repo
// 	entry_repo     Entry_repo
// }

// /*
// - if we need to transfer 10 usd from acc 1 to acc 2
// 	=> check that the acc 1 has 10 usd
// 	=> create a transfer record with amount = 10 for these 2 accounts (from & to)
// 	=> create an acc entry record for acc 1 with amount = -10
// 	=> create an acc entry record for acc 2 with amount = +10
// 	=> update the balance of acc 1 by subtracting 10 usd
// 	=> update the balance of acc 2 by adding 10 usd
// */

// // the transaction function
// func (ts *TransactionStore) TransferMoneyTx(tx_args TransferTxParams) (*TransferTxResult, error) {
// 	var tx_result TransferTxResult
// 	var update_acc_params Update_account_params
// 	// var qry_result interface{}

// 	// define a context
// 	ctx, cancel := CreateContext()
// 	defer cancel()

// 	// start the transaciton
// 	tx, err := ts.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		log.Println("Error while beginning the transaction : ", err)
// 		return nil, err
// 	}

// 	// execute queries within the transaction
// 	//! 1. check that the sending acc has 10 usd
// 	// fetch the account
// 	balance, err := ts.acc_repo.GetBalanceById(tx_args.FromAccountID)
// 	if *balance < tx_args.Amount {
// 		log.Println("The balance of the accoutn is less than the required amount for the transaction")
// 		// rollback the transaction
// 		tx.Rollback()
// 		// return the error
// 		return nil, err
// 	}
// 	// acc_balance, err := tx.QueryRowContext(ctx, `SELECT * FROM accounts WHERE id = $1`, tx_args.FromAccountID)
// 	if err != nil {
// 		log.Println("Error while executing the query that checks that the sending account has the specified balance for the money transfer transaction : ", err)
// 		// rollback the transaction
// 		tx.Rollback()
// 		// return the error
// 		return nil, err
// 	}

// 	//! 2. create a transfer record with amount = 10 for these 2 accounts (from & to)
// 	transfer_record := domain.Transfer{
// 		FromAccountID: int64(tx_args.FromAccountID),
// 		ToAccountID:   int64(tx_args.ToAccountID),
// 		Amount:        float64(tx_args.Amount),
// 	}
// 	err = ts.trans_repo.Create(&transfer_record)
// 	if err != nil {
// 		log.Println("Error while executing the query that Create a new transfer record in the database for the money transfer transaction : ", err)
// 		// rollback the transaction
// 		tx.Rollback()
// 		// return the error
// 		return nil, err
// 	}
// 	// then set the persisted data to the result
// 	tx_result.Transfer, err = ts.trans_repo.GetByFromAndToAccountId(tx_args.FromAccountID, tx_args.ToAccountID)
// 	if err != nil {
// 		log.Println("Error while retrieving the persisted transfer record after executing the query that Create a new transfer record in the database for the money transfer transaction : ", err)
// 		// rollback the transaction
// 		tx.Rollback()
// 		// return the error
// 		return nil, err
// 	}

// 	//! 3. create an acc entry record for acc 1 with amount = -10
// 	entry_record_of_the_from_account := domain.Entry{AccountID: int64(tx_args.FromAccountID), Amount: float64(tx_args.Amount * -1)}
// 	err = ts.entry_repo.Create(&entry_record_of_the_from_account)
// 	if err != nil {
// 		log.Println("Error while creating a new entry record for the account which transfers the money for the money transfer transaction : ", err)
// 		// rollback the transaction
// 		tx.Rollback()
// 		// return the error
// 		return nil, err
// 	}
// 	// then set the persisted data to the result
// 	tx_result.FromEntry, err = ts.entry_repo.GetByAccountId(tx_args.FromAccountID)
// 	if err != nil {
// 		log.Println("Error while retrieving the persisted entry record of the account which transfers the money for the money transfer transaction : ", err)
// 		// rollback the transaction
// 		tx.Rollback()
// 		// return the error
// 		return nil, err
// 	}

// 	//! 4. create an acc entry record for acc 2 with amount = +10
// 	entry_record_of_the_to_account := domain.Entry{AccountID: int64(tx_args.ToAccountID), Amount: float64(tx_args.Amount)}
// 	err = ts.entry_repo.Create(&entry_record_of_the_to_account)
// 	if err != nil {
// 		log.Println("Error while creating a new entry record for the account which recieve the money for the money transfer transaction : ", err)
// 		// rollback the transaction
// 		tx.Rollback()
// 		// return the error
// 		return nil, err
// 	}
// 	// then set the persisted data to the result
// 	tx_result.ToEntry, err = ts.entry_repo.GetByAccountId(tx_args.ToAccountID)
// 	if err != nil {
// 		log.Println("Error while retrieving the persisted entry record of the account which recieve the money for the money transfer transaction : ", err)
// 		// rollback the transaction
// 		tx.Rollback()
// 		// return the error
// 		return nil, err
// 	}
// 	//! 5. update the balance of acc 1 by subtracting 10 usd
// 	updated_balance_of_from_account := tx_args.Amount * float64(-1)
// 	update_acc_params.Balance = &updated_balance_of_from_account
// 	updated_from_acc, err := ts.acc_repo.Update(tx_args.FromAccountID, update_acc_params)
// 	if err != nil {
// 		log.Println("Error while retrieving the persisted entry record of the account which transfer the money for the money transfer transaction : ", err)
// 		// rollback the transaction
// 		tx.Rollback()
// 		// return the error
// 		return nil, err
// 	}
// 	tx_result.FromAccount = updated_from_acc

// 	//! 6. update the balance of acc 2 by adding 10 usd
// 	updated_balance_of_to_account := tx_args.Amount * float64(-1)
// 	update_acc_params.Balance = &updated_balance_of_to_account
// 	updated_to_acc, err := ts.acc_repo.Update(tx_args.ToAccountID, update_acc_params)
// 	if err != nil {
// 		log.Println("Error while retrieving the persisted entry record of the account which recieve the money for the money transfer transaction : ", err)
// 		// rollback the transaction
// 		tx.Rollback()
// 		// return the error
// 		return nil, err
// 	}
// 	tx_result.FromAccount = updated_to_acc

// 	// commit the transaction
// 	err = tx.Commit()
// 	if err != nil {
// 		log.Println("Error while commiting the transaction : ", err)
// 		return nil, err
// 	}

// 	log.Println("Transfer Moeny Transaction is succeed!")

// 	return &tx_result, nil

// }
