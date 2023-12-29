package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

func CreateClient(url string) *sql.DB {
	db, err := sql.Open("libsql", url)

	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	create, err := os.ReadFile("resources/sql/create.sql")

	if err != nil {
		log.Fatal("Error reading create.sql:", err)
	}

	_, err = db.Exec(string(create))

	if err != nil {
		log.Fatal("Error creating database:", err)
	}

	return db
}
