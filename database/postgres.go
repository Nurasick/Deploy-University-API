package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

func OpenConnection(dbURL string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return conn
}
func OpenConnectionPool(dbURL string) *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), dbURL)
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

func CloseConnectionPool(pool *pgxpool.Pool) {
	pool.Close()
}

func PGXConnToSQLDB(conn *pgx.Conn) *sql.DB {
	sqlDB := sql.OpenDB(stdlib.GetConnector(*conn.Config()))

	return sqlDB
}

type DBConnWrapper struct {
	SQLDB *sql.DB
}

func NewDBConnWrapper(sqlDB *sql.DB) *DBConnWrapper {
	return &DBConnWrapper{SQLDB: sqlDB}
}

func (db *DBConnWrapper) Ping(ctx context.Context) error {
	return db.SQLDB.PingContext(ctx)
}

func PingPool(pool *pgxpool.Pool) error {
	return pool.Ping(context.Background())
}
