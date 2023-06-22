package config

import (
	"os"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func GetConfig() *Config {
	return &Config{
		Username: getEnv("POSTGRES_USER", "postgres"),
		Password: getEnv("POSTGRES_PASSWORD", "postgres"),
		Host:     getEnv("POSTGRES_HOST", "localhost"),
		Port:     getEnv("POSTGRES_PORT", "5432"),
		DBName:   getEnv("POSTGRES_DB", "postgres"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
