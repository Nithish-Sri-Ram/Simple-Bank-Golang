package main

import (
	"database/sql"
	"log"

	"github.com/Nithish-Sri-Ram/simplebank/api"
	db "github.com/Nithish-Sri-Ram/simplebank/db/sqlc"
	"github.com/Nithish-Sri-Ram/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	// The above returns a connection object in return
	if err != nil {
		log.Fatal("cannot connect to the DB:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
