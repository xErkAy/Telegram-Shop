package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net"
)

func main() {

	if initBot() != nil {
		fmt.Println("Error occurred while connecting to the bot")
		return
	}
	if initDB() != nil {
		fmt.Println("Error occurred while connecting to the database")
		return
	}
	go socketHandler()

	defer database.Action.Close()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.Action.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			go CallbackQueryHandler(update)
		} else if update.Message != nil {
			go UserHandler(update)
		}
	}
}

func socketHandler() {
	listener, _ := net.Listen("tcp", "172.20.10.3:8001")
	for {
		connection, err := listener.Accept()
		if err == nil {
			go handleSocketConnection(connection)
		}
	}
}
