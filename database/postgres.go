package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

func OpenConnection(dbURL string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return conn
}

func CloseConnection(conn *pgx.Conn) {
	if err := conn.Close(context.Background()); err != nil {
		log.Println("Error closing DB connection:", err)
	}
}

func PGXConnToSQLDB(conn *pgx.Conn) *sql.DB {
	sqlDB := sql.OpenDB(stdlib.GetConnector(*conn.Config()))

	return sqlDB
}
