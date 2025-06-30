package config

import (
	"log"
	"os"
)

type Config struct {
	Telegram struct {
		Token string
	}
	Application struct {
		Port string
	}
	Database struct {
		DSN string
	}
	Vk struct {
		BaseUrl     string
		ApiVersion  string
		AccessToken string
	}
}

func Init() *Config {
	return &Config{
		Telegram: struct {
			Token string
		}{
			Token: getEnvByKey("TELEGRAM_API_TOKEN", ""),
		},
		Application: struct {
			Port string
		}{
			Port: getEnvByKey("APP_PORT", "8080"),
		},
		Database: struct {
			DSN string
		}{
			DSN: getEnvByKey("DB_DSN", "postgres://forge:secret@postgres:5432/app?sslmode=disable&search_path=public"),
		},
		Vk: struct {
			BaseUrl     string
			ApiVersion  string
			AccessToken string
		}{
			BaseUrl:     getEnvByKey("VK_BASE_URL", "https://api.vk.com"),
			ApiVersion:  getEnvByKey("VK_API_VERSION", "5.199"),
			AccessToken: getEnvByKey("VK_ACCESS_TOKEN", ""),
		}}
}

func getEnvByKey(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue == "" {
			log.Fatalf("Missing required environment variable %s", key)
		}

		value = defaultValue
	}

	return value
}
