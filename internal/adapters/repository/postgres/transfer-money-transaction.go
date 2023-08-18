package postgres

// import (
// 	"gobanking/internal/core/domain"
// )

// /*
// transactions
// - if we need to transfer 10 usd from acc 1 to acc 2
// 	=> check that the acc 1 has 10 usd
// 	=> create a transfer record with amount = 10 for these 2 accounts (from & to)
// 	=> create an acc entry record for acc 1 with amount = -10
// 	=> create an acc entry record for acc 2 with amount = +10
// 	=> update the balance of acc 1 by subtracting 10 usd
// 	=> update the balance of acc 2 by adding 10 usd
// */

// // TransferTxParams is the required fields to perform money transaction between 2 accounts
// type TransferTxParams struct {
// 	FromAccountID int     `json:"from_account_id"`
// 	ToAccountID   int     `json:"to_account_id"`
// 	Amount        float64 `json:"amount"`
// }

// // TransferTxResult is the result of our transaction
// type TransferTxResult struct {
// 	// the created Transfer domain entity
// 	Transfer *domain.Transfer `json:"transfer"`

// 	// the account that transfers the money after updating its balance
// 	FromAccount *domain.Account `json:"from_account"`

// 	// the accoutn that got the money after updating its balance
// 	ToAccount *domain.Account `json:"to_account"`

// 	// the entry record that stores the money transfer from the transferring account
// 	FromEntry *domain.Entry `json:"from_entry"`

// 	// the entry record that stores the money which transfered to the reciever account
// 	ToEntry *domain.Entry `json:"to_entry"`
// }
