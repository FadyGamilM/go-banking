package types

import (
	"gobanking/internal/core-layer/domain"
	"time"
)

type CreateTransferRequest struct {
	ToAccountID   int64   `json:"to_account_id"`
	FromAccountID int64   `json:"from_account_id"`
	Amount        float64 `json:"currency"`
}

type CreateTransferResponse struct {
	ID            int64     `json:"id"`
	ToAccountID   int64     `json:"to_account_id"`
	FromAccountID int64     `json:"from_account_id"`
	Amount        float64   `json:"currency"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type GetTransferResponse struct {
	ID            int64     `json:"id"`
	ToAccountID   int64     `json:"to_account_id"`
	FromAccountID int64     `json:"from_account_id"`
	Amount        float64   `json:"currency"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TransferMoneyTransactionParam struct {
	ToAccountID   int64   `json:"to_account_id"`
	FromAccountID int64   `json:"from_account_id"`
	Amount        float64 `json:"currency"`
}

type TransferMoneyTransactionResult struct {
	// created transfer record
	Transfer *domain.Transfer `json:"transfer"`
	// From account after it's balance is updated
	FromAccount *domain.Account `json:"from_account"`
	// To account after it's balance is updated
	ToAccount *domain.Account `json:"to_account"`
	// The entry of the from account (records moving out money)
	FromEntry *domain.Entry `json:"from_entry"`
	// The entry of the to account (records moving in money)
	ToEntry *domain.Entry `json:"to_entry"`
}
