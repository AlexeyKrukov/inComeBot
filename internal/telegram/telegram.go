package telegram

import (
	"log"
	"strconv"

	"github.com/AlexeyKrukov/inComeBot/internal/incomes"
	"github.com/AlexeyKrukov/inComeBot/internal/pkg/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	bot *tgbotapi.BotAPI
}

type config interface {
	IsDebug() string
	TelegramToken() string
}

func New(cfg config) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken())

	if err != nil {
		return nil, err
	}

	bot.Debug, _ = strconv.ParseBool(cfg.IsDebug())

	return &Telegram{
		bot: bot,
	}, nil
}

func (t *Telegram) Run() {
	log.Printf("Authorized on account %s", t.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	data := storage.New()

	for update := range updates {
		if update.Message != nil {

			incomes.CalculateIncomes(update, data)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Income added")

			_, err := t.bot.Send(msg)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (t *Telegram) Shutdown() {
	t.bot.StopReceivingUpdates()
}
