package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(value string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env" + err.Error())
	}

	return os.Getenv(value)
}
