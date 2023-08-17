package types

import "time"

type CreateTransferRequest struct {
	OwnerName string  `json:"owner_name"`
	Balance   float64 `json:"balance"`
	Currency  string  `json:"currency"`
}

type CreateTransferResponse struct {
	ID        int64     `json:"id"`
	OwnerName string    `json:"owner_name"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetTransferResponse struct {
	ID        int64     `json:"id"`
	OwnerName string    `json:"owner_name"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
