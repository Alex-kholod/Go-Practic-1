package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_DSN     string
	BcryptCost int // например, 12
	Addr       string
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, errors.New("error loading env variables")
	}

	cost := 12
	if v := os.Getenv("BCRYPT_COST"); v != "" {
		// необязательно: распарсить int, при ошибке оставить 12
	}
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode,
	)
	if dsn == "" {
		return Config{}, errors.New("DB_DSN is empty")
	}

	return Config{
		DB_DSN:     dsn,
		BcryptCost: cost,
		Addr:       addr,
	}, nil
}
