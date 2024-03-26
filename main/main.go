package main

import (
	"os"
	"telegram_bot/functions"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAMBOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	functions.MessageEcho(bot)
}
