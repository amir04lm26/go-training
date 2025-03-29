package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost     = "DB_HOST"
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
