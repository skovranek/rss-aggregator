package database

import (
	"database/sql"
	"log"
)

type CloseDB func() error

// remember: defer db.Close()
func InitDB(dbURL string) (*Queries, CloseDB) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error: initDB(dbURL): sql.Open(): cannot open database: %v", err)
	}

	return New(db), db.Close
}
