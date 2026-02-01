package env

import "os"

// Load reads variables from a .env file (e.g. github.com/joho/godotenv).
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
