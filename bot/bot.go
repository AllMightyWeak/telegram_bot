package bot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAMBOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	return bot
}
