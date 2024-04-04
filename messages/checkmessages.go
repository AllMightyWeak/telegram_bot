package messages

import (
	"log"
	"telegram_bot/connection"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CheckForMessages(telegramBot *tgbotapi.BotAPI) {
	bot := telegramBot

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			CheckCommandType(bot, &update)
		}
	}
}

func CheckCommandType(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите /qrt для подключения кошелька телеграм\nВведите /qrk для подключения кошелька Tonkeeper")
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	case "qrt":
		walletName := "Wallet"
		connection.SetUpConnection(bot, update.Message, walletName)

	case "qrk":
		walletName := "Tonkeeper"
		connection.SetUpConnection(bot, update.Message, walletName)
	}

}
