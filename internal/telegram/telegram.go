package telegram

import (
	"awesomeProject1/internal/config"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"sync"
)

var incomes = make(map[string]int)
var mx sync.Mutex

type Telegram struct {
	bot *tgbotapi.BotAPI
}

func New(cfg *config.Config) *Telegram {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)

	if err != nil {
		log.Fatalf("Telegram API error: %s", err)
	}

	bot.Debug = cfg.IsDebug

	return &Telegram{
		bot: bot,
	}
}

func (t *Telegram) Run() {
	log.Printf("Authorized on account %s", t.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			income, _ := strconv.Atoi(update.Message.Text)

			mx.Lock()

			incomes[update.Message.From.UserName] = incomes[update.Message.From.UserName] + income

			mx.Unlock()

			fmt.Printf("%+v\n", incomes)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Income added")

			t.bot.Send(msg)
		}
	}
}
