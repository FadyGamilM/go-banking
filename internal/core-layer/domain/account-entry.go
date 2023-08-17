package domain

import "time"

type AccountEntry struct {
	ID        int64
	AccountID int64
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}