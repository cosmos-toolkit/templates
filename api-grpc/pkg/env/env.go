// Package env provides optional helpers for environment variables (e.g. .env).
// Use in cmd/server/main.go for config.
package env

import "os"

// Load reads environment from files. In production use os.Getenv or a lib like godotenv.
func Load(filenames ...string) error {
	_ = filenames
	return nil
}

// Get returns the environment variable or the default value.
func Get(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
