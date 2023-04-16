package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rhiadc/gobank/cmd/api"
	"github.com/rhiadc/gobank/config"
	db "github.com/rhiadc/gobank/db/sqlc"
)

func main() {
	env := config.LoadEnvVars()
	conn, err := sql.Open(env.Database.Driver, env.Database.GetURIConnection())
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	serverAddr := fmt.Sprintf("localhost:%s", env.APIPort)
	store := db.NewStore(conn)
	server := api.NewServer(store)
	if err := server.Start(serverAddr); err != nil {
		log.Fatal("Error starting API server:", err)
	}
}
