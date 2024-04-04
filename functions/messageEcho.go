package functions

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageEcho(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	if _, err := bot.Send(msg); err != nil {
		log.Printf("ERROR: %s", err)
		msg = tgbotapi.NewMessage(message.Chat.ID, "Кажется что-то пошло не так! Я не могу ответить на ваше сообщение")
		bot.Send(msg)
	}
}
