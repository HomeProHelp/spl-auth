package db

import (
	"fmt"
	"github/LissaiDev/spl-auth/pkg/hermes"
	"github/LissaiDev/spl-auth/utils"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		utils.GetEnv("DB_HOST", "localhost"),
		utils.GetEnv("POSTGRES_USER", "spl-auth"),
		utils.GetEnv("POSTGRES_PASSWORD", "spl-auth"),
		utils.GetEnv("POSTGRES_DB", "spl-auth"),
		utils.GetEnv("DB_PORT", "5432"),
	)

	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		hermes.Log(3, fmt.Sprintf("Error connecting to database: %s", err), true)
	}

	hermes.Log(1, "Database connection was successfull", false)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
