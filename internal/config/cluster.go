package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DSN() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("failed to load .env file")
	}

	pgDsn := os.Getenv("POSTGRES_PG_DSN")
	sslMode := os.Getenv("POSTGRES_SSL_MODE")

	return fmt.Sprintf("%s&sslmode=%s", pgDsn, sslMode)
}

type Cluster struct {
	Conn *pgx.Conn
}

func NewCluster(ctx context.Context) (*Cluster, error) {
	dsn := DSN()

	fmt.Printf("dsn: %s\n", dsn)
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect: %w", err)
	}

	return &Cluster{Conn: conn}, nil
}
