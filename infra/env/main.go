package env

import (
	"fmt"
	"os"
)

func GetUrl() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	return fmt.Sprintf("0.0.0.0:%s", port)
}

func GetDatabaseUrl() string {
	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		databaseUrl = "file:dev.sqlite3?_journal_mode=WAL&mode=rwc&cache=shared"
	}

	return databaseUrl
}
