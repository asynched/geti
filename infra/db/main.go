package db

import (
	"database/sql"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

func CreateClient(url string) *sql.DB {
	db, err := sql.Open("libsql", url)

	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	return db
}
