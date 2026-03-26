package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/momoyo-droid/capim/api/internal/utils"
)

type Config struct {
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func LoadConfig() (*Config, error) {
	godotenv.Load()

	port := os.Getenv("PORT")
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	if DBHost == "" {
		return nil, utils.ErrDBHostRequired
	}

	if DBUser == "" {
		return nil, utils.ErrDBUserRequired
	}

	if DBPassword == "" {
		return nil, utils.ErrDBPasswordRequired
	}

	if DBName == "" {
		return nil, utils.ErrDBNameRequired
	}

	if DBPort == "" {
		return nil, utils.ErrDBPortRequired
	}
	if _, err := strconv.Atoi(DBPort); err != nil {
		return nil, utils.ErrInvalidDBPort
	}

	if port == "" {
		return nil, utils.ErrPortRequired
	}

	if _, err := strconv.Atoi(port); err != nil {
		return nil, utils.ErrInvalidPort
	}

	cfg := &Config{
		Port:       port,
		DBHost:     DBHost,
		DBUser:     DBUser,
		DBPassword: DBPassword,
		DBName:     DBName,
		DBPort:     DBPort,
	}

	return cfg, nil
}
