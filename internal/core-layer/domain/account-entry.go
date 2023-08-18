package domain

import "time"

type Entry struct {
	ID        int64
	AccountID int64
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
