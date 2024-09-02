package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    var err error
    dsn := "user=postgres password=admin dbname=api sslmode=disable"
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)

    if err := createTables(); err != nil {
        log.Fatalf("Failed to create tables: %v", err)
    }
}

func createTables() error {
    const query = `
    CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime TIMESTAMP NOT NULL,
        user_id INTEGER
    )`
    _, err := DB.Exec(query)
    return err
}
