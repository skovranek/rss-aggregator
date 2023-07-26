package database

import (
    "context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB    *database.Queries
	Limit int32
}

func TestDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(`Error: database_test.go: godotenv.Load(): cannot load ".env" file`)
	}
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error: database_test.go: sql.Open(): cannot open database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	cfg := apiConfig{
		DB:    dbQueries,
		Limit: int32(10),
	}

    // TODO: 
    // add copy of .env into here, and add to ignore on github
    // test database funcs

}
