package env

import (
	"github.com/deidelson/go-chi-api/pkg/core/convertion"
	"os"
)

func GetEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

func GetEnvOrDefaultAsInt(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		parsedValue, err := convertion.StringToInt(value)
		if err != nil {
			panic("Cannot load env variable")
		}
		return parsedValue
	}
	return defaultValue
}