package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the bot.
type Config struct {
	IsDebug       bool   `mapstructure:"IS_DEBUG"`
	TelegramToken string `mapstructure:"TELEGRAM_TOKEN"`
	MigrationURL  string `mapstructure:"MIGRATION_URL"`
	DBSource      string `mapstructure:"DB_SOURCE"`
}

func (c *Config) GetDebug() bool {
	return c.IsDebug
}

func (c *Config) GetTelegramToken() string {
	return c.TelegramToken
}

func New(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
