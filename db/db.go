package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return fmt.Errorf("variável DATABASE_URL não configurada")
	}

	var err error
	Conn, err = pgx.Connect(ctx, url)
	if err != nil {
		return fmt.Errorf("falha ao conectar no PostgreSQL: %w", err)
	}

	fmt.Println("✅ Banco de dados conectado com sucesso!")
	return nil
}
