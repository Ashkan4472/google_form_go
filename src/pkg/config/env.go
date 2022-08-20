package config

import (
	"github.com/joho/godotenv"
)

func InitialEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
