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
		is_active, order_id := isActiveChat(chat)
		if is_active {
			formData := url.Values{
				"user_id":      {chat.user_id},
				"order_id":     {strconv.Itoa(order_id)},
				"message_text": {update.Message.Text},
				"is_sender":    {"True"},
			}
			http.PostForm("http://localhost:8000/api/messages/", formData)
			return
		}

		response, _ := db.Query("SELECT * FROM shop_users WHERE user_id=$1", chat.user_id)
		if isUserNew(response, chat) {
			SendMessage(chat.ID, "Добро пожаловать, "+chat.first_name+"!")
		} else {
			go updateUserInfo(response, chat)
		}
		go SendKeyboard(chat.ID)
	}
}

func CallbackQueryHandler(update tgbotapi.Update) {
	if update.CallbackQuery.Data == "makeorder" {
		go SendMessage(update.CallbackQuery.Message.Chat.ID, "Временно недоступно.")
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
