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
		Host       string
		Username   string
		Password   string
		Port       string
		Database   string
		DisableSsl string
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
			Host       string
			Username   string
			Password   string
			Port       string
			Database   string
			DisableSsl string
		}{
			Host:       getEnvByKey("DB_HOST", "postgres"),
			Username:   getEnvByKey("DB_USERNAME", "forge"),
			Password:   getEnvByKey("DB_PASSWORD", "secret"),
			Port:       getEnvByKey("DB_PORT", "5432"),
			Database:   getEnvByKey("DB_DATABASE", "app"),
			DisableSsl: getEnvByKey("DB_SSL_MODE", "disable"),
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
			log.Fatalf("Missing required environment variable %s", value)
		}

		value = defaultValue
	}

	return value
}
