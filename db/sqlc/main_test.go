package db

import (
	"database/sql"
	"testing"
	"os"
	"log"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

var (
	dbDriver = os.Getenv("DB_DRIVER")
	dbSource = os.Getenv("BB_SOURCE")
)
func TestMain(m *testing.M) {

	var err error 
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("could not connect to the database", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())

}