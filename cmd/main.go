package main

import (
	"context"
	"fmt"

	"github.com/AlexeyKrukov/inComeBot/internal/config"
	"github.com/AlexeyKrukov/inComeBot/internal/telegram"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	var tg telegram.Telegram
	ctx := context.Background()
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		cfg := config.New()

		tg, err := telegram.New(cfg)

		if err != nil {
			log.Fatal().Err(err).Msg("Error while starting")
		}

		tg.Run()

		return nil
	})

	g.Go(func() error {
		<-gCtx.Done()

		tg.Shutdown()
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(errors.Wrap(err, "exit reason"))
	}
}
