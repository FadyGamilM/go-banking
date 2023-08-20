package types

import "time"

type CreateTransferRequest struct {
	OwnerName string  `json:"owner_name"`
	Balance   float64 `json:"balance"`
	Currency  string  `json:"currency"`
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
