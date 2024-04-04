package main

import (
	"telegram_bot/bot"
	"telegram_bot/messages"
)

func main() {
	telegrambot := bot.StartBot()
	messages.CheckForMessages(telegrambot)
}
