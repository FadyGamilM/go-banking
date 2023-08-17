package domain

import "time"

type Transfer struct {
	ID            int64
	ToAccountID   int64
	FromAccountID int64
	Amount        float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
