package config


import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

func loadEnvVariables() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func GetEnv(key string, defaultVal ...string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	if len(defaultVal) > 0 {
		return defaultVal[0]
	}

	return ""
}
