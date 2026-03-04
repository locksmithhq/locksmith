package database

import (
	"context"
	"fmt"
	"os"

	"github.com/booscaaa/initializers/postgres"
	"github.com/booscaaa/initializers/postgres/types"
)

var conn types.Database

func Initialize(ctx context.Context) {
	if os.Getenv("SCHEMA") == "" {
		os.Setenv("SCHEMA", "locksmith")
	}

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("SSL_MODE"),
	)

	conn = postgres.Initialize(
		ctx,
		databaseURL,
		postgres.WithMigrations(false),
	)

	_, err := conn.ExecContext(ctx, fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", os.Getenv("SCHEMA")))
	if err != nil {
		panic(err)
	}

	databaseURL = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("SCHEMA"),
		os.Getenv("SSL_MODE"),
	)
	conn = postgres.Initialize(
		ctx,
		databaseURL,
	)
}

func GetConnection() types.Database {
	return conn
}
