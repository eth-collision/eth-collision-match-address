package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func sendMsgText(text string) {
	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Println(err)
	}
	message := tgbotapi.NewMessage(tgChatId, text)
	_, err = bot.Send(message)
	if err != nil {
		log.Println(err)
	}
}
