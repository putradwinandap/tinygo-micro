package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	RabbitMQURL string
	RedisURL    string
}

func LoadConfig() *Config {
	return &Config{
		Port:        LoadEnv("PORT", "8080"),
		DatabaseURL: LoadEnv("DATABASE_URL", "postgres://user:password@postgres:5432/tinygo_micro?sslmode=disable"),
		RabbitMQURL: LoadEnv("RABBITMQ_URL", "amqp://user:password@rabbitmq:5672/"),
		RedisURL:    LoadEnv("REDIS_URL", "redis://redis:6379/0"),
	}
}

func LoadEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
