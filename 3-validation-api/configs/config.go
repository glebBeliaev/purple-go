package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Mail MailConfig
}

type MailConfig struct {
	SMTP     string
	Address  string
	Password string
}

func LoadConfig() *Config {
	err := godotenv.Load(filepath.Join("3-validation-api", ".env"))
	if err != nil {
		log.Println("Error loading .env file, using defaults")
	}
	return &Config{
		Mail: MailConfig{
			SMTP:     os.Getenv("EMAIL_SMTP"),
			Address:  os.Getenv("EMAIL_ADDRESS"),
			Password: os.Getenv("EMAIL_PASSWORD"),
		},
	}
}
