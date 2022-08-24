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
		is_chat_active, order_id := isActiveChat(chat.ID)
		if is_chat_active {
			formData := url.Values{
				"user_id":      {chat.user_id},
				"order_id":     {strconv.Itoa(order_id)},
				"message_text": {update.Message.Text},
				"is_sender":    {"True"},
			}
			http.PostForm("http://localhost:8000/api/messages/", formData)
			return
		}

		if isOrderActive(chat.ID) {
			makeNewOrder(chat.ID, update.Message.Text)
			return
		}

		is_new, err := isUserNew(chat)
		if err != nil {
			SendMessage(chat.ID, "Произошла ошибка. Попробуйте еще раз.")
			return
		}
		if is_new {
			SendMessage(chat.ID, "Добро пожаловать, "+chat.first_name+"!")
		}
		go SendKeyboard(chat.ID)
	}
}

func CallbackQueryHandler(update tgbotapi.Update) {
	if update.CallbackQuery.Data == "makeorder" {
		user_id := update.CallbackQuery.Message.Chat.ID
		is_chat_active, _ := isActiveChat(user_id)
		if is_chat_active {
			go SendMessage(user_id, "У вас активный чат! Завершите его, чтобы сделать новый заказ.")
			return
		}

		if isOrderActive(user_id) {
			go SendMessage(user_id, "Завершите предыдущий заказ, чтобы сделать новый.")
			return
		}

		go makeOrderActive(user_id)

	} else if update.CallbackQuery.Data == "getmenu" {
		go SendDocument(update.CallbackQuery.Message.Chat.ID, "menu.pdf")
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

	json_err := json.Unmarshal(bytes.Trim(data, "'\x00'"), &chat)
	if json_err != nil {
		fmt.Println(json_err)
		return
	}

	SendMessage(chat.User_id, chat.Message)
	connection.Close()
}
