package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

// this interface type is used as a dependency in all repos so any repo can receive a transaction "tx" or the database pool itself "db" and execute the transaction via anyone of them because both implements DBTX
type DBTX interface {
	// i named the params here because i need the users to know what they should pass to this func
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)

	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// is the wrapper that all repos will depends on
type PG struct {
	DB DBTX
}

func NewPG(db DBTX) *PG {
	return &PG{DB: db}
}

func (pg *PG) WithTx(tx *sql.Tx) *PG {
	return &PG{
		DB: tx,
	}
}

// ==> For individual queries | For connection to the database
type AppDB struct {
	// private access to the `db` package only
	db_pool *sql.DB
}

func Connect(dsn string) (*sql.DB, error) {
	dbArgs := dbArgs{
		DbTimeOut:     3 * time.Second,
		maxOpenDbConn: 10,
		maxIdleDbConn: 5,
		maxDbLifeTime: 5 * time.Minute,
	}

	conn_pool, err := connectToPostgresInstance(dsn, dbArgs)
	if err != nil {
		return nil, err
	}

	return conn_pool, nil
}

func (db *AppDB) Close() {
	db.db_pool.Close()
}

func (db *AppDB) GetDbPool() *sql.DB {
	return db.db_pool
}
