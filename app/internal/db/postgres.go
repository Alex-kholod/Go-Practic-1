package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	// currentDir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal("failed to get current directory: %w", err)
	// }
	// pathToEnv := filepath.Join(currentDir, "..", "..", "..", ".env")
	// godotenv.Load(pathToEnv)

	dbHost := getEnvValue("DB_HOST")
	dbUser := getEnvValue("DB_USER")
	dbPassword := getEnvValue("DB_PASSWORD")
	dbName := getEnvValue("DB_NAME")
	dbPort := getEnvValue("DB_PORT")
	dbSSLMode := getEnvValue("DB_SSL_MODE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode,
	)
	if dsn == "" {
		log.Fatal("DB_DSN is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal("connect db:", err)
	}

	// Настроим пул соединений
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func getEnvValue(key string) string {
	_ = godotenv.Load("C:\\EDU\\Go\\pz6\\.env")
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("env value is empty")
	}
	return value
}
