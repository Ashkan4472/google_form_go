package config

import (
	"os"

	"github.com/Ashkan4472/google_form_go/src/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialDatabase() {
	var err error
	dsn := os.Getenv("DATABASE_CONNECTION_STR")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.User{})
}
