package db

import (
	"fmt"
	"github/LissaiDev/spl-auth/pkg/hermes"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDatabaseInstance() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		GetEnv("DB_HOST", "localhost"),
		GetEnv("POSTGRES_USER", "spl-auth"),
		GetEnv("POSTGRES_PASSWORD", "spl-auth"),
		GetEnv("POSTGRES_DB", "spl-auth"),
		GetEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		hermes.Log(3, fmt.Sprintf("Error connecting to database: %s", err), true)
	}

	hermes.Log(1, "Database connection was successfull", false)
	return db
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetEnv(key, fallback string) string {
	if env, exists := os.LookupEnv(key); exists {
		return env
	}
	return fallback
}
