package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port      string        `validate:"required"`
	JWTTTL    time.Duration `validate:"required"`
	SecretKey string        `validate:"required"`
	DB        DBConfig      `validate:"required"`
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	URL      string // For DATABASE_URL when provided
}

func (cfg *DBConfig) String() string {
	// If URL is set (from DATABASE_URL env var), return it directly
	if cfg.URL != "" {
		return cfg.URL
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}

// GetConfig loads env vars and returns AppConfig
func GetConfig() *AppConfig {
	// Load .env if exists (local dev)
	_ = godotenv.Load(".env") // optional: only for local dev

	ttlStr := os.Getenv("JWT_TTL")
	if ttlStr == "" {
		ttlStr = "24h" // default 1 day
	}
	jwtTTL, err := time.ParseDuration(ttlStr)
	if err != nil {
		log.Fatalf("Invalid JWT_TTL: %v", err)
	}

	dbPort := 5432
	if portStr := os.Getenv("POSTGRES_PORT"); portStr != "" {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatalf("Invalid POSTGRES_PORT: %v", err)
		}
		dbPort = p
	}

	dbCfg := DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     dbPort,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		URL:      os.Getenv("DATABASE_URL"), // Use DATABASE_URL if provided (e.g., Render)
	} // Use POSTGRES_DB, not POSTGRES_NAME
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		cfg := &AppConfig{
			Port:      os.Getenv("PORT"),
			JWTTTL:    jwtTTL,
			SecretKey: os.Getenv("JWT_SECRET"),
			DB:        dbCfg,
		}
		return cfg
	}

	cfg := &AppConfig{
		Port:      os.Getenv("PORT"),
		JWTTTL:    jwtTTL,
		SecretKey: os.Getenv("JWT_SECRET"),
		DB:        dbCfg,
	}

	// Validate required fields
	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		log.Fatalf("Config validation failed: %v", err)
	}

	return cfg
}
