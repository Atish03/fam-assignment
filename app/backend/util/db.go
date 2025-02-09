package util

import (
	"database/sql"
	"fmt"
	"os"
	"log"
)

var DB *sql.DB

func InitDB() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	init_file := os.Getenv("INIT_FILE")

	var err error

	// Connecting to db
	connStr := fmt.Sprintf("postgres://%s:%s@postgres:5432/%s?sslmode=disable", user, password, db)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	query, err := os.ReadFile(init_file)
	if err != nil {
		log.Fatal("Error reading SQL file:", err)
	}

	// Executing init.sql
	_, err = DB.Exec(string(query))
	if err != nil {
		log.Fatal("Error executing SQL:", err)
	}

	fmt.Println("Database initialized successfully!")
}