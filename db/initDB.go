// db/initDB.go

package db

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitializeDB initializes the SQLite database connection and creates the necessary table schema
func InitializeDB(dbPath string, schemaFile string) *sql.DB {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection to the database is successful
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")

	// Read SQL schema from file
	schemaSQL, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		log.Fatal("Failed to read SQL schema file:", err)
	}

	// Execute SQL schema to create tables
	if _, err := db.Exec(string(schemaSQL)); err != nil {
		log.Fatal("Failed to create database tables:", err)
	}

	log.Println("Database schema created")

	return db
}
