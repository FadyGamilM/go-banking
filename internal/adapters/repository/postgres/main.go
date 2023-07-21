package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"
)

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
func ConnectToPostgresInstance(DSN string, args DbArgs) (*sql.DB, error) {
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

func CreateContext() (context.Context, context.CancelFunc) {
	// define context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	return ctx, cancel
}
