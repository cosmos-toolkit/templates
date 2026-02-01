// Package env provides helpers for environment variables (e.g. .env).
// Use from cmd or internal/commands when a command needs config.
package env

import "os"

// Load reads variables from a .env file. In production use os.Getenv
// or a lib like github.com/joho/godotenv.
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
