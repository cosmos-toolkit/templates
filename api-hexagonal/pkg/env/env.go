// Package env é um plugin opcional para carregar variáveis de ambiente (ex.: .env).
// Uso: config, err := env.Load() no cmd/api/main.go
package env

import "os"

// Load lê variáveis do ambiente. Em produção use os.Getenv diretamente
// ou uma lib como github.com/joho/godotenv para .env.
func Load(filenames ...string) error {
	// Exemplo: return godotenv.Load(filenames...)
	_ = filenames
	return nil
}

// Get retorna a variável de ambiente ou o valor default.
func Get(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
