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
	db   *sql.DB   // The global DB connection instance
	once sync.Once // Ensures DB initialization happens only once
	err  error     // To store initialization error
)

// InitDB initializes the database connection (Singleton Pattern).
func InitDB() (*sql.DB, error) {
	once.Do(func() {
		// Load environment variables from .env file
		err = godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found, relying on environment variables")
		}

		// Retrieve DB URL and Token from environment variables
		dbURL := os.Getenv("TURSO_DATABASE_URL")
		dbTOKEN := os.Getenv("TURSO_AUTH_TOKEN")

		// Check if necessary credentials are missing
		if dbURL == "" || dbTOKEN == "" {
			err = fmt.Errorf("database credentials missing")
			return
		}

		// Prepare the connection URL with the authentication token
		url := dbURL + "?authToken=" + dbTOKEN

		// Open a new DB connection
		db, err = sql.Open("libsql", url)
		if err != nil {
			err = fmt.Errorf("failed to open database: %v", err)
			return
		}

		// Test the connection
		err = db.Ping()
		if err != nil {
			err = fmt.Errorf("failed to connect to database: %v", err)
			return
		}

		log.Println("Database connected successfully")
	})

	return db, err
}

// GetDB returns the existing database connection. It must be initialized before calling this.
func GetDB() *sql.DB {
	return db
}

// CloseDB closes the database connection.
func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error closing the database connection:", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}
