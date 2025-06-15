package config

import "os"

// Config содержит конфигурацию приложения
type Config struct {
	Port         string
	DatabasePath string
	JWTSecret    string
	GinMode      string
}

// New создает новую конфигурацию
func New() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		DatabasePath: getEnv("DB_PATH", "./database.db"),
		JWTSecret:    getEnv("JWT_SECRET", "default-secret-key"),
		GinMode:      getEnv("GIN_MODE", "debug"),
	}
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 