package postgres

import (
	"context"
	"time"
)

func CreateContext() (context.Context, context.CancelFunc) {
	// define context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	return ctx, cancel
}
