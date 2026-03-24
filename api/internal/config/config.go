package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/momoyo-droid/capim/api/internal/utils"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		return nil, utils.ErrPortRequired
	}

	if _, err := strconv.Atoi(port); err != nil {
		return nil, utils.ErrInvalidPort
	}

	cfg := &Config{
		Port: port,
	}

	return cfg, nil
}
