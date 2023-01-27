package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug  bool `yaml:"is-debug" env-default:"false"`
	Telegram struct {
		Token string `yaml:"token" env-required:"true"`
	} `yaml:"telegram"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig("../config.yaml", instance); err != nil {
			help, err := cleanenv.GetDescription(instance, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
