package functions

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageEcho(telegramBot *tgbotapi.BotAPI) {
	bot := telegramBot

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if _, err := bot.Send(msg); err != nil {
			log.Printf("ERROR: %s", err)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Кажется что-то пошло не так! Я не могу ответить на ваше сообщение")
			bot.Send(msg)
		}
	}
}
