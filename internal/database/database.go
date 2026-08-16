package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Connect() (*pgx.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	databaseURL := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}