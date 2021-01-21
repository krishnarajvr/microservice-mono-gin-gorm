package common

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Print("Not able to get current working director")
	}
	// loads values from .env into the system
	if err := godotenv.Load(dir + "/.env"); err != nil {
		fmt.Print("No .env file found")
	}
}

// GetEnv returns an environment variable or a default value if not present
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}

// LoadEnvVars will load a ".env[.development|.test]" file if it exists and set ENV vars.
// Useful in development and test modes. Not used in production.
func LoadEnvVars() {
	env := GetEnv("GIN_ENV", "development")

	if env == "production" || env == "staging" {
		log.Println("Not using .env file in production or staging.")
		return
	}

	filename := ".env." + env

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = ".env"
	}

	err := godotenv.Load(filename)
	if err != nil {
		log.Println(".env file not loaded")
	}
}
