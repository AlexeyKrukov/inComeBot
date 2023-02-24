package config

import (
	"os"
)

type conf struct {
	isDebug       string
	telegramToken string
}

type Config struct {
	conf conf
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		conf: conf{
			isDebug:       getEnv("IS_DEBUG", "false"),
			telegramToken: getEnv("TELEGRAM_TOKEN", "XXXX"),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func (c *Config) IsDebug() string {
	return c.conf.isDebug
}

func (c *Config) TelegramToken() string {
	return c.conf.telegramToken
}
