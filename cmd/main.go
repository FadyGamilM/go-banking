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
	"github.com/FadyGamilM/go-banking-v2/utils"
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

	configs, err := utils.LoadConfig("..")
	log.Println("the db host", configs.DbHost)
	if err != nil {
		log.Fatalf("error while loading config files => %v", err)
	}

	// get the *sql.DB and *db.PG instance
	_, PG := setupDatabaseConnection(configs)

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

	server.Start(configs.ServerAddr)

}

func setupDatabaseConnection(configs utils.DevConfig) (*sql.DB, *db.PG) {
	var err error

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", configs.DbHost, configs.DbPort, configs.DbUsername, configs.DbPassword, configs.DbName)

	connPool, err := db.Connect(connStr)
	if err != nil {
		log.Fatalf("error trying to connect to test db : %v \n", err)
	}
	pg := db.NewPG(connPool)

	return connPool, pg
}
