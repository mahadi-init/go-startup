package api

import (
	"gin-app/db"
	"log"
)

func ConnectDb() {
	// Initialize DB connection
	_, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %s", err)
	}

	// Defer closing the DB connection
	defer db.CloseDB()

	// Example query
	database := db.GetDB()
	var version string
	err = database.QueryRow("SELECT sqlite_version()").Scan(&version)
	if err != nil {
		log.Fatalf("Query failed: %s", err)
	}

	log.Println("SQLite Version:", version)
}
