package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/FadyGamilM/go-banking-v2/api"
	"github.com/FadyGamilM/go-banking-v2/db"
	"github.com/FadyGamilM/go-banking-v2/internal/account"
	"github.com/FadyGamilM/go-banking-v2/internal/entry"
	"github.com/FadyGamilM/go-banking-v2/internal/transfer"
)

const (
	// Replace these values with your PostgreSQL container settings
	host       = "127.0.0.1" // Use the container IP if needed
	port       = 5555        // Default PostgreSQL port
	user       = "dev"       // PostgreSQL username
	password   = "dev"       // PostgreSQL password
	dbname     = "devdb"     // Database name
	serverAddr = "0.0.0.0:8000"
)

func main() {

	// get the *sql.DB and *db.PG instance
	_, PG := setupDatabaseConnection()

	log.Println("setup db ")
	// get all repos
	accountRepo := account.NewAccountRepo(PG)
	entryRepo := entry.NewEntryRepo(PG)
	transferRepo := transfer.NewTransferRepo(PG)

	accountService := account.NewAccountService(accountRepo)
	entryService := entry.NewEntryService(entryRepo)
	transferService := transfer.NewTransferService(transferRepo)

	accountHandler := api.NewAccountHandler(accountService)
	entryHandler := api.NewEntryHandler(entryService)
	transferHandler := api.NewTransferHandler(transferService)

	server := api.NewServer(accountHandler, entryHandler, transferHandler)

	server.Start(serverAddr)

}

func setupDatabaseConnection() (*sql.DB, *db.PG) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error

	connPool, err := db.Connect(connStr)
	if err != nil {
		log.Fatalf("error trying to connect to test db : %v \n", err)
	}
	pg := db.NewPG(connPool)

	return connPool, pg
}
