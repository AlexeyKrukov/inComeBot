package incomes

import (
	"strconv"

	storagePackage "github.com/AlexeyKrukov/inComeBot/internal/pkg/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CalculateIncomes(update tgbotapi.Update, data *storagePackage.Storage) {
	income, _ := strconv.ParseFloat(update.Message.Text, 32)

	data.GetMutex().Lock()

	data.SetBalance(update.Message.From.UserName, income)

	data.GetMutex().Unlock()
}
