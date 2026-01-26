package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect(connString string) {
	var err error
	Conn, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("DB connection failed", err)
	}
}
