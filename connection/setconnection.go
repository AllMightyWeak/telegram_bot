package connection

import (
	"context"
	"fmt"
	"log"
	"os"
	"telegram_bot/generator"
	"telegram_bot/responses"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/exp/maps"

	tonconnect "github.com/cameo-engineering/tonconnect"
)

func SetUpConnection(bot *tgbotapi.BotAPI, message *tgbotapi.Message, walletName string) {
	for _, wallet := range tonconnect.Wallets {
		switch wallet.Name {
		case walletName:

			s, connreq := ConnectToSession()

			//DeepLinkGenerate(bot, s, message, connreq)
			universalLink := UniversalLinkGenerate(bot, message, s, connreq, wallet)
			//ConnectToWallet(bot, s, message)

			msg := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FilePath("qr.png"))

			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonURL("link", universalLink),
				),
			)

			msg.ReplyMarkup = keyboard
			bot.Send(msg)
			os.Remove("qr.png")
		}

	}
}

func DeepLinkGenerate(bot *tgbotapi.BotAPI, s *tonconnect.Session, message *tgbotapi.Message, connreq *tonconnect.ConnectRequest) {
	deepLink, err := s.GenerateDeeplink(*connreq)
	if err != nil {
		log.Println("\nMANIFEST READ ERROR")
		log.Println(err)
	}
	wrappedDeepLink := tonconnect.WrapDeeplink(deepLink)

	generator.GenerateQRCode(wrappedDeepLink)
	responses.PhotoResponce(bot, message)
	os.Remove("qr.png")

	responses.LinkResponce(bot, message, wrappedDeepLink)
}

func UniversalLinkGenerate(bot *tgbotapi.BotAPI, message *tgbotapi.Message, s *tonconnect.Session, connreq *tonconnect.ConnectRequest, wallet tonconnect.Wallet) string {
	link, err := s.GenerateUniversalLink(wallet, *connreq)
	if err != nil {
		log.Println("\nMANIFEST READ ERROR")
		log.Println(err)
	}
	generator.GenerateQRCode(link)
	//responses.PhotoResponce(bot, message)
	//os.Remove("qr.png")
	//responses.LinkResponce(bot, message, link)
	return link

}

func ConnectToSession() (*tonconnect.Session, *tonconnect.ConnectRequest) {
	s, err := tonconnect.NewSession()
	if err != nil {
		log.Println("\nSESSION ERROR")

		log.Fatal(err)
	}

	connreq, err := tonconnect.NewConnectRequest(
		"https://raw.githubusercontent.com/AllMightyWeak/AllMightyWeak.github.io/main/tonconnect-mainfest.json")
	if err != nil {
		log.Println("\nMANIFEST CONNECT ERROR")
		log.Println(err)
	}

	return s, connreq
}

func ConnectToWallet(bot *tgbotapi.BotAPI, s *tonconnect.Session, message *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	res, err := s.Connect(ctx, (maps.Values(tonconnect.Wallets))...)
	if err != nil {
		log.Fatal(err)
	}

	var addr string
	network := "mainnet"
	for _, item := range res.Items {
		if item.Name == "ton_addr" {
			addr = item.Address
			if item.Network == -3 {
				network = "testnet"
			}
		}
	}
	responses.TextResponce(bot, message, "Connected app "+res.Device.AppName)
	fmt.Printf(
		"%s %s for %s is connected to %s with %s address\n\n",
		res.Device.AppName,
		res.Device.AppVersion,
		res.Device.Platform,
		network,
		addr,
	)
}
