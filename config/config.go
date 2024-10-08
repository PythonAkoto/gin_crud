package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config structure to hold all configuration values
type Config struct {
	AppEnv     string
	Port       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

// LoadConfig reads configuration from environment variables or .env file
func LoadConfig() Config {
	// Load .env file (optional, you can remove this if you don't want to use godotenv)
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found, reading from environment variables.")
	}

	config := Config{
		AppEnv:     getEnv("APP_ENV", "development"),
		Port:       getEnv("PORT", "8080"),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbUser:     getEnv("DB_USER", ""),
		DbPassword: getEnv("DB_PASSWORD", ""),
		DbName:     getEnv("DB_NAME", ""),
	}

	return config
}

// Helper function to get environment variables with a fallback default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
