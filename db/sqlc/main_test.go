package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/SaifAlqady51/simple-bank/util"
	_ "github.com/lib/pq"
)

var testQuery *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot read config ", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("we can't connect to the database")
	}

	testQuery = New(testDB)
	os.Exit(m.Run())
}
