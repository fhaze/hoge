package internal

import "os"

func GetEnvDefault(name string, def string) string {
	if val := os.Getenv(name); val != "" {
		return val
	}
	return def
}

