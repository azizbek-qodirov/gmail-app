package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	GATEWAY_HTTP_PORT string
	GMAIL_GRPC_PORT   string
	SENDER_EMAIL      string
	APP_PASSWORD      string
	REDIS_URL         string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.REDIS_URL = cast.ToString(coalesce("REDIS_URL", "localhost:6379"))
	config.GMAIL_GRPC_PORT = cast.ToString(coalesce("GMAIL_GRPC_PORT", ":50050"))
	config.GATEWAY_HTTP_PORT = cast.ToString(coalesce("GATEWAY_HTTP_PORT", ":8088"))
	config.SENDER_EMAIL = cast.ToString(coalesce("SENDER_EMAIL", "email@example.com"))
	config.APP_PASSWORD = cast.ToString(coalesce("APP_PASSWORD", "your_app_password_here"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
