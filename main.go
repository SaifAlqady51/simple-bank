package main

import (
	"database/sql"
	"log"

	"github.com/SaifAlqady51/simple-bank/api"
	db "github.com/SaifAlqady51/simple-bank/db/sqlc"
	"github.com/SaifAlqady51/simple-bank/util"
	_ "github.com/lib/pq"
)

func main() {
	// get env vairables
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config ", err)
	}
	// connect to the database
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("we can't connect to the database")
	}
	// new Store
	store := db.NewStore(conn)
	// start a new Server
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start a server ", err)
	}
}
