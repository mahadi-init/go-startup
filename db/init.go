package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql" // Import the Turso libsql driver
)

var (
	db   *sql.DB
	once sync.Once
)

// InitDB initializes the database connection.
func InitDB() (*sql.DB, error) {
	var err error
	once.Do(func() {
		err = godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found, relying on environment variables")
		}

		dbURL := os.Getenv("TURSO_DATABASE_URL")
		dbTOKEN := os.Getenv("TURSO_AUTH_TOKEN")

		if dbURL == "" || dbTOKEN == "" {
			err = fmt.Errorf("database credentials missing")
			return
		}

		url := dbURL + "?authToken=" + dbTOKEN

		db, err = sql.Open("libsql", url)
		if err != nil {
			log.Fatalf("failed to open database: %s", err)
		}

		// Test the connection
		err = db.Ping()
		if err != nil {
			log.Fatalf("failed to connect to database: %s", err)
		}

		log.Println("Database connected successfully")
	})

	return db, err
}

// GetDB returns the existing database connection.
func GetDB() *sql.DB {
	return db
}

// CloseDB closes the database connection.
func CloseDB() {
	if db != nil {
		db.Close()
		log.Println("Database connection closed")
	}
}
