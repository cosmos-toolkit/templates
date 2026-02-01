// Package env provides helpers for environment variables in Lambda.
// Lambda injects config via os.Getenv (e.g. from function configuration or .env in SAM).
package env

import "os"

// Get returns the environment variable or the default value.
func Get(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
