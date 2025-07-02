package database

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/lib/pq"
	"github.com/Shashanktriathi1703/student-api/internal/config"
)

func NewPostgresConnection(cfg *config.Config) *sql.DB {
	// Create connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error Opening database %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database %v", err)
	}

	//Create users table if not exists
	createTableQuery :=
		`CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY, 
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL, 
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`

	// Excueting the query wiothout returning the rows.
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating users table : %v", err)
	}

	return db
}
