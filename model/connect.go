package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var con *sql.DB

func Connect() (*sql.DB, error) {
	connStr := "user=postgres dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
	con = db
	return db, nil
}
