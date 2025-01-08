package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	// we dont actually call any function which uses the above package but we need it for the db connection
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

// This is a special function - by convention the test main funcion is the main entry point of all unit tests inside 1 specific golang package

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	// The above returns a connection object in return
	if err != nil {
		log.Fatal("cannot connect to the DB:", err)
	}

	testQueries = New(conn)
	// Else we use the connection to create a test querie object

	os.Exit(m.Run())
	// This above function will run the unit test and it will return the exit code which tell us wheter the tests pass or fail - if fail it will simply
}
