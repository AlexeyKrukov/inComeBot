package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config is a struct for configs
type Config struct {
	IsDebug  bool `yaml:"is-debug" env-default:"false"`
	Telegram struct {
		Token string `yaml:"token" env-required:"true"`
	} `yaml:"telegram"`
}

func (c *Config) GetDebug() bool {
	return c.IsDebug
}

func (c *Config) GetTelegramToken() string {
	return c.Telegram.Token
}

func New() (*Config, error) {
	instance := &Config{}

	//спросить про конфиги
	if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
		help, _ := cleanenv.GetDescription(instance, nil)

		return nil, errors.New(help)

	}

	return instance, nil
}
