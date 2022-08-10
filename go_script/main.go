package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot, err = tgbotapi.NewBotAPI("5288667624:AAFBARM244Me22WXxZiTbK1r8kTfqAgM01I")

func main() {
	go checkErrors(err)
	go checkErrors(db_err)
	defer db.Close()

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		go UserHandler(update)
	}
}

func checkErrors(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
