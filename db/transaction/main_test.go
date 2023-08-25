package transaction

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/FadyGamilM/go-banking-v2/db"
	account_repo "github.com/FadyGamilM/go-banking-v2/internal/account"
	account "github.com/FadyGamilM/go-banking-v2/internal/account/domain"
)

var TestPG *db.PG
var TestSqlDB *sql.DB
var accTestRepo account.AccountRepo

// var test_acc_repo *account.AccountRepo

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

	connPool, err := db.Connect(connStr)
	if err != nil {
		log.Fatalf("error trying to connect to test db : %v \n", err)
	}
	testPG := db.NewPG(connPool)
	TestSqlDB = connPool

	TestPG = testPG

	test_acc_repo := account_repo.NewAccountRepo(testPG)
	accTestRepo = test_acc_repo

	os.Exit(m.Run())
}

func createAccountForTest(ctx context.Context, req_args *account.Account) (*account.Account, error) {

	created_acc, err := accTestRepo.Create(ctx, req_args)

	return created_acc, err
}
