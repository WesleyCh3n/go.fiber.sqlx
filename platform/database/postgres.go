package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func PostgreSQLConn() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}
	return db, err
}
