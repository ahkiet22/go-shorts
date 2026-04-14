package database

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

// singleton pattern for database connection pool
var (
	instance *pgxpool.Pool
	once     sync.Once
)

func GetDatabasePool() *pgxpool.Pool {
	once.Do(func() {
		conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal("Cannot connect DB:", err)
		}
		instance = conn
	})

	log.Println("Database connection pool initialized")

	return instance
}
