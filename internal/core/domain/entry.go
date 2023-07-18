package domain

import "time"

type entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// the change of the amount in a bank account can be by +ve or -ve
	Amount    float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
