package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JWTSecret []byte
	JWTTTL    time.Duration
}

func Load() Config {
	_ = godotenv.Load()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is required")
	}
	ttl := os.Getenv("JWT_TTL")
	if ttl == "" {
		ttl = "24h"
	}
	dur, err := time.ParseDuration(ttl)
	if err != nil {
		log.Fatal("bad JWT_TTL")
	}

	return Config{Port: ":" + port, JWTSecret: []byte(secret), JWTTTL: dur}
}
