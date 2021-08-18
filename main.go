package main

import (
	"database/sql"
	"log"

	"github.com/gaku101/my-portfolio/api"
	db "github.com/gaku101/my-portfolio/db/sqlc"
	_ "github.com/lib/pq"
)

const  (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/my_portfolio?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
