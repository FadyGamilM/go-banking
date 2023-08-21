package postgres

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

type PG_DB struct {
	DB *sql.DB
}

type PG_TX struct {
	TX *sql.Tx
}

type dbArgs struct {
	DbTimeOut     time.Duration
	maxOpenDbConn int
	maxIdleDbConn int
	maxDbLifeTime time.Duration
}

// factory method to create a new Postgres database wrapper
func NewPgDb(dsn string) (*PG_DB, error) {
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

	return &PG_DB{DB: conn_pool}, nil
}

// function to test the connection before applying the connection
func testDB(db *sql.DB) error {
	// ping the database instance first and check if there is a response or an error returned
	err := db.Ping()
	if err != nil {
		log.Printf("cannot ping the postgres instance \n ERROR ➜ %v", err)
		return err
	}

	// if no error, we recieved the pong
	log.Println("we ping, the postgres instance pong successfully :D")
	return nil // return no errors
}

// function to perform the actuall connection
func connectToPostgresInstance(DSN string, args dbArgs) (*sql.DB, error) {
	// open a pool of connections
	pool_of_conn, err := sql.Open("pgx", DSN)
	if err != nil {
		log.Printf("cannot open a connection \n ERROR ➜ %v", err)
		return nil, err
	}
	// set pool conn attributes
	pool_of_conn.SetMaxOpenConns(args.maxOpenDbConn)
	pool_of_conn.SetMaxIdleConns(args.maxIdleDbConn)
	pool_of_conn.SetConnMaxLifetime(args.maxDbLifeTime)

	// ping to the connection
	ping_conn_err := testDB(pool_of_conn)
	if ping_conn_err != nil {
		return nil, ping_conn_err
	}

	// return the response
	return pool_of_conn, nil
}
