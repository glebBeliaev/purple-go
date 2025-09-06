package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Token string
}

func LoadConfig() *Config {
	err := godotenv.Load("learning/.env")
	if err != nil {
		log.Println("Error loading .env file, using defaults")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Token: os.Getenv("TOKEN"),
		},
	}
}
