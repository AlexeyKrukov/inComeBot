package telegram

import (
	"github.com/AlexeyKrukov/inComeBot/internal/pkg/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

type Telegram struct {
	bot     *tgbotapi.BotAPI
	storage storage.Storage
}

type Config interface {
	GetDebug() bool
	GetTelegramToken() string
}

func New(cfg Config) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.GetTelegramToken())

	if err != nil {
		return nil, err
	}

	bot.Debug = cfg.GetDebug()

	return &Telegram{
		bot: bot,
	}, nil
}

func (t *Telegram) Run() {
	log.Printf("Authorized on account %s", t.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	t.storage.Incomes = make(storage.IncomesById)

	for update := range updates {
		if update.Message != nil {

			income, _ := strconv.ParseFloat(update.Message.Text, 2)

			t.storage.Mu.Lock()

			t.storage.Incomes[update.Message.From.UserName] = t.storage.Incomes[update.Message.From.UserName] + income

			t.storage.Mu.Unlock()

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Income added")

			t.bot.Send(msg)
		}
	}
}

func (t *Telegram) Shutdown() {
	t.bot.StopReceivingUpdates()
}
