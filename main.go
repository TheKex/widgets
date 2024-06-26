package main

import (
	"database/sql"
	"log"
	"widgets/api"
	db "widgets/db/sqlc"
	"widgets/util"

	_ "github.com/lib/pq"
	_ "go.uber.org/mock/mockgen/model"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("fail to load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
