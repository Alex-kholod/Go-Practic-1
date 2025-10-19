package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	if err := loadEnv(); err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

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

func loadEnv() error {
	// Получаем путь к текущему файлу (postgres.go)
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get current file path")
	}

	// Вычисляем путь к .env относительно расположения postgres.go
	currentDir := filepath.Dir(filename)
	projectRoot := filepath.Join(currentDir, "..", "..", "..")
	envPath := filepath.Join(projectRoot, ".env")

	// Нормализуем путь для Windows
	envPath = filepath.Clean(envPath)

	log.Printf("Looking for .env at: %s", envPath)

	// Проверяем существует ли файл
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		return fmt.Errorf(".env file not found at: %s", envPath)
	}

	// Загружаем .env файл
	if err := godotenv.Load(envPath); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	log.Printf("✅ .env loaded from: %s", envPath)
	return nil
}

func getEnvValue(key string) string {
	//_ = godotenv.Load("C:\\EDU\\Go\\pz6\\.env")
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("env value is empty")
	}
	return value
}
