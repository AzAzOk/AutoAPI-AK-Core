package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port            string
	PostgresDSN     string
	RedisAddr       string
	RedisPassword   string
	VaultMasterKey  string
	AccountVaultURL string // URL сервиса account-vault
}

func Load() *Config {
	return &Config{
		Port:            getEnv("ROUTER_PORT", "8080"),
		PostgresDSN:     getEnv("POSTGRES_DSN", ""),
		RedisAddr:       getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:   getEnv("REDIS_PASSWORD", ""),
		VaultMasterKey:  getEnv("VAULT_MASTER_KEY", ""),
		AccountVaultURL: getEnv("ACCOUNT_VAULT_URL", "http://account-vault:3002"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}
