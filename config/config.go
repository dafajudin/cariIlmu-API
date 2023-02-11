package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/tkanos/gonfig"
)

//receive configuration from .env file
type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func GetConfig(key string) string {
	
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	
	return os.Getenv(key)
}

