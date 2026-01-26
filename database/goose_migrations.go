package database

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

func GooseMigrate(sqlDB *sql.DB, migrationDir string) {
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Goose set dialect error: %v", err)
	}

	if err := goose.Up(sqlDB, migrationDir); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Println("Database migrations applied successfully")
}
