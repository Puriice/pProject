package env

import (
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	godotenv.Load()
}

func GetEnv(key string, fallback string) string {
	env, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return env
}
