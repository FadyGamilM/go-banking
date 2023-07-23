package postgres

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

type Postgres_repo struct {
	db *sql.DB
}

type DbArgs struct {
	DbTimeOut     time.Duration
	maxOpenDbConn int
	maxIdleDbConn int
	maxDbLifeTime time.Duration
}

// factory pattern
func NewPostgresRepo(dsn string) (*Postgres_repo, error) {
	db_args := DbArgs{
		DbTimeOut:     3 * time.Second,
		maxOpenDbConn: 10,
		maxIdleDbConn: 5,
		maxDbLifeTime: 5 * time.Minute,
	}

	conn_pool, err := ConnectToPostgresInstance(dsn, db_args)
	if err != nil {
		return nil, err
	}

	return &Postgres_repo{db: conn_pool}, nil
}
