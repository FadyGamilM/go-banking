package postgres

import (
	"context"
	"fmt"
	"gobanking/internal/infra-layer/db/postgres"
	"log"
	"os"
	"testing"
)

// create global instance of the PG_DB object which is used to execute queries
var test_PG_DB *postgres.PG_DB
var test_PG_TX *postgres.PG_TX
var test_acc_repo *PG_AccountRepository
var test_entry_repo *PG_EntryRepository
var test_transfer_Repo *PG_TransferRepository

// func init() {
// 	test_PG_DB = *postgres.PG_DB{}
// }

func TestMain(m *testing.M) {
	// initiate

	// Replace these values with your PostgreSQL container settings
	host := "127.0.0.1" // Use the container IP if needed
	port := 3333        // Default PostgreSQL port
	user := "test"      // PostgreSQL username
	password := "test"  // PostgreSQL password
	dbname := "testdb"  // Database name

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	test_PG_DB, err = postgres.NewPgDb(connStr)
	if err != nil {
		log.Fatalf("error trying to connect to test db : %v \n", err)
	}

	err = test_PG_DB.DB.Ping()
	if err != nil {
		log.Fatalf("error trying to ping to the test db : %v \n", err)
	}

	// instantiate a new transaction
	tx, err := test_PG_DB.DB.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatalf("error trying to begin a transaction : %v \n", err)
	}
	test_PG_TX = &postgres.PG_TX{TX: tx}
	test_acc_repo = NewPG_AccountRepo(test_PG_TX)
	test_entry_repo = NewPG_EntryRepo(test_PG_TX)
	test_transfer_Repo = NewPG_TransferRepo(test_PG_TX)

	os.Exit(m.Run())
}
