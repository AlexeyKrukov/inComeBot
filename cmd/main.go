package main

import (
	"context"
	"fmt"
	"github.com/AlexeyKrukov/inComeBot/internal/config"
	"github.com/AlexeyKrukov/inComeBot/internal/telegram"
	"github.com/rs/zerolog/log"
	"os/signal"
	"syscall"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	cfg, err := config.New(".")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	tg, err := telegram.New(&cfg)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot create instance")
	}

	go tg.Run()

	<-ctx.Done()

	tg.Shutdown()

	fmt.Println("shutting down server gracefully")
}
