package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

func UserHandler(update tgbotapi.Update) {
	chat := Chat{update.Message.Chat.ID, update.Message.MessageID, strconv.FormatInt(update.Message.From.ID, 10), string(update.Message.From.UserName), string(update.Message.From.FirstName)}

	switch update.Message.Command() {
	default:
		isChatActive, orderId := database.isActiveChat(chat.ID)
		if isChatActive {
			formData := url.Values{
				"user_id":      {chat.userId},
				"order_id":     {strconv.Itoa(orderId)},
				"message_text": {update.Message.Text},
				"is_sender":    {"True"},
			}
			http.PostForm("http://localhost:8000/api/messages/", formData)
			return
		}

		if database.isOrderActive(chat.ID) {
			database.makeNewOrder(chat.ID, update.Message.Text)
			return
		}

		isNew, err := database.isUserNew(chat)
		if err != nil {
			bot.SendMessage(chat.ID, "Произошла ошибка. Попробуйте еще раз.")
			return
		}
		if isNew {
			bot.SendMessage(chat.ID, "Добро пожаловать, "+chat.firstName+"!")
		}
		go bot.SendKeyboard(chat.ID)
	}
}

func CallbackQueryHandler(update tgbotapi.Update) {
	if update.CallbackQuery.Data == "makeorder" {
		userId := update.CallbackQuery.Message.Chat.ID
		isChatActive, _ := database.isActiveChat(userId)
		if isChatActive {
			go bot.SendMessage(userId, "У вас активный чат! Завершите его, чтобы сделать новый заказ.")
			return
		}

		if database.isOrderActive(userId) {
			go bot.SendMessage(userId, "Завершите предыдущий заказ, чтобы сделать новый.")
			return
		}

		go database.makeOrderActive(userId)

	} else if update.CallbackQuery.Data == "getmenu" {
		go bot.SendDocument(update.CallbackQuery.Message.Chat.ID, "menu.pdf")
	} else if update.CallbackQuery.Data == "getordersstatus" {
		go SendOrdersStatus(update.CallbackQuery.From.ID)
	}
}

func handleSocketConnection(connection net.Conn) {
	var chat Notify

	data := make([]byte, 4096)
	_, err := connection.Read(data)
	if err != nil {
		return
	}

	jsonErr := json.Unmarshal(bytes.Trim(data, "'\x00'"), &chat)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return
	}

	bot.SendMessage(chat.UserId, chat.Message)
	connection.Close()
}
