package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Mail MailConfig
}

type MailConfig struct {
	SMTP     string
	Adress   string
	Password string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using defaults")
	}
	return &Config{
		Mail: MailConfig{
			SMTP:     os.Getenv("EMAIL_SMTP"),
			Adress:   os.Getenv("EMAIL_ADRESS"),
			Password: os.Getenv("EMAIL_PASSWORD"),
		},
	}
}
