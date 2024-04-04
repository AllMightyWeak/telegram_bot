package responses

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func PhotoResponce(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FilePath("qr.png"))
	if _, err := bot.Send(msg); err != nil {
		log.Printf("ERROR: %s", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Кажется что-то пошло не так! Я не могу ответить на ваше сообщение")
		bot.Send(msg)
	}
}

func TextResponce(bot *tgbotapi.BotAPI, message *tgbotapi.Message, text string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	if _, err := bot.Send(msg); err != nil {
		log.Printf("ERROR: %s", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Кажется что-то пошло не так! Я не могу ответить на ваше сообщение")
		bot.Send(msg)
	}
}

func StartCommandResponce(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, Bot!")

	if _, err := bot.Send(msg); err != nil {
		log.Printf("ERROR: %s", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Кажется что-то пошло не так! Я не могу ответить на ваше сообщение")
		bot.Send(msg)
	}
}

func DefaultCommandResponce(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Странная команда, попробуй ещё раз")

	if _, err := bot.Send(msg); err != nil {
		log.Printf("ERROR: %s", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Кажется что-то пошло не так! Я не могу ответить на ваше сообщение")
		bot.Send(msg)
	}
}

func LinkResponce(bot *tgbotapi.BotAPI, message *tgbotapi.Message, link string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, link)

	if _, err := bot.Send(msg); err != nil {
		log.Printf("ERROR: %s", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Кажется что-то пошло не так! Я не могу ответить на ваше сообщение")
		bot.Send(msg)
	}
}
