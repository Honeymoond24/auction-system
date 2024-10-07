package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseURI string
}

func NewConfig() *Config {
	_ = godotenv.Load()
	databaseURI := os.Getenv("DATABASE_URI")

	if databaseURI == "" {
		log.Fatal("Environment variable is missing")
	}

	return &Config{
		DatabaseURI: databaseURI,
	}
}
