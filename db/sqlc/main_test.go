package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Nithish-Sri-Ram/simplebank/util"
	_ "github.com/lib/pq"
	// we dont actually call any function which uses the above package but we need it for the db connection
)

var testQueries *Queries
var testDB *sql.DB

// This is a special function - by convention the test main funcion is the main entry point of all unit tests inside 1 specific golang package

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	// The above returns a connection object in return
	if err != nil {
		log.Fatal("cannot connect to the DB:", err)
	}

	testQueries = New(testDB)
	// Else we use the connection to create a test querie object

	os.Exit(m.Run())
	// This above function will run the unit test and it will return the exit code which tell us wheter the tests pass or fail - if fail it will simply
}
