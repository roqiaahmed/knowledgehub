package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(envPath string) (*gorm.DB, error) {
	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("error loading .env file from %s: %w", envPath, err)
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		return nil, fmt.Errorf("Database credentials are not fully set in the environment variables")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database!")
	}
	return db, nil
}
