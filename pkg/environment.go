package pkg

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var REQUIRED_ENV_KEYS = []string{"POSTGRES_USER", "POSTGRES_PASSWORD",
	"POSTGRES_PORT", "POSTGRES_HOST", "POSTGRES_DB", "GOOGLE_CLIENT_ID",
	"GOOGLE_CLIENT_SECRET", "GOOGLE_REDIRECT_URI"}

func LoadEnvironment(paths ...string) error {
	err := godotenv.Load(paths...)
	if err != nil {
		return err
	}

	for _, v := range REQUIRED_ENV_KEYS {
		k := os.Getenv(v)
		if k == "" {
			return fmt.Errorf("no required .env variable '%v'", v)
		}
	}

	return nil
}

func GetEnvironment(key string, fallback interface{}) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	m, ok := fallback.(string)
	if !ok {
		return ""
	}
	return m
}
