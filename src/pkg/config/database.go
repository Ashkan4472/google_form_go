package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitialDatabase() {
	dsn := os.Getenv("DATABASE_CONNECTION_STR")
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
