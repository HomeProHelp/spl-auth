package utils

import "os"

func GetEnv(key, fallback string) string {
	if env, exists := os.LookupEnv(key); exists {
		return env
	}
	return fallback
}
