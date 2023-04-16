package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	db          *sql.DB
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/gobank_test?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	testQueries = New(db)
	os.Exit(m.Run())
}
