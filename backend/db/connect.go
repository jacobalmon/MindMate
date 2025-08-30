package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool is the global connection pool.
var Pool *pgxpool.Pool

func Connect() {
	url := os.Getenv("DATABASE_URL")
	var err error
	Pool, err = pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Connected to database.")
}
