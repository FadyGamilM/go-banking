package domain

import "time"

type account struct {
	ID        int64     `json:"id"`
	OwnerName string    `json:"owner_name"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
