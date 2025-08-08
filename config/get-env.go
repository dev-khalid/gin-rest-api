package config

import (
	"os"

	"github.com/joho/godotenv"
)

func loadEnvVariables() {
	godotenv.Load()
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
