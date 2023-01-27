package main

import (
	"awesomeProject1/internal/config"
	"awesomeProject1/internal/telegram"
)

func main() {
	cfg := config.GetConfig()

	tg := telegram.New(cfg)

	tg.Run()
}
