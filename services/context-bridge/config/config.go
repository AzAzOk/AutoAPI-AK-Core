package config

import "os"

type Config struct {
	Port        string
	PostgresDSN string
	RedisAddr   string
	RedisPassword string
	// Максимум токенов в контексте по умолчанию (если провайдер не указал)
	DefaultMaxTokens int
}

func Load() *Config {
	return &Config{
		Port:             getEnv("CONTEXT_BRIDGE_PORT", "8081"),
		PostgresDSN:      getEnv("POSTGRES_DSN", ""),
		RedisAddr:        getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		DefaultMaxTokens: 4096,
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
