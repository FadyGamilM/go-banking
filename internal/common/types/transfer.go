package types

import "time"

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
<<<<<<< HEAD
}

type TransferMoneyTransactionParam struct {
	ToAccountID   int64   `json:"to_account_id"`
	FromAccountID int64   `json:"from_account_id"`
	Amount        float64 `json:"currency"`
}

type TransferMoneyTransactionResult struct {
	ToAccountID   int64   `json:"to_account_id"`
	FromAccountID int64   `json:"from_account_id"`
	Amount        float64 `json:"currency"`
=======
>>>>>>> b358ac80580b9f2601913429f7976c3b093b3efa
}
