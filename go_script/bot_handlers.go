package main

import (
	"database/sql"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

var db, db_err = sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=shop_database sslmode=disable")
var status = [3]string{"не взят в работу", "готовится", "готов к выдаче"}

func UserHandler(update tgbotapi.Update) {
	// if !update.Message.IsCommand() {
	// 	return
	// }

	if update.CallbackQuery != nil {
		if update.CallbackQuery.Data == "makeorder" {
			go SendMessage(update.CallbackQuery.Message.Chat.ID, "Временно недоступно.")
		} else if update.CallbackQuery.Data == "getmenu" {
			go SendDocument(update.CallbackQuery.Message.Chat.ID, "menu.pdf")
		} else if update.CallbackQuery.Data == "getordersstatus" {
			response, _ := db.Query("SELECT order_id, status FROM shop_orders WHERE user_id_id=$1 and is_closed=FALSE", update.CallbackQuery.From.ID)
			go SendOrdersStatus(response, update.CallbackQuery.From.ID)
		}
		return
	}

	chat := Chat{update.Message.Chat.ID, update.Message.MessageID, strconv.FormatInt(update.Message.From.ID, 10), string(update.Message.From.UserName), string(update.Message.From.FirstName)}

	switch update.Message.Command() {
	default:
		response, _ := db.Query("SELECT * FROM shop_users WHERE user_id = $1", chat.user_id)
		if isUserNew(response, chat) {
			SendMessage(chat.ID, "Добро пожаловать, "+chat.first_name+"!")
		} else {
			go updateUserInfo(response, chat)
		}
		go SendKeyboard(chat.ID)
	}
}

func isUserNew(response *sql.Rows, chat Chat) bool {
	err_count := 0
	if !response.Next() {
		for {
			_, err := db.Exec("INSERT INTO shop_users VALUES ($1, $2, $3)", chat.user_id, chat.user_name, chat.first_name)
			if err == nil {
				break
			} else {
				err_count += 1
			}
			if err_count == 10 {
				break
			}
		}
		return true
	}
	return false
}

func updateUserInfo(response *sql.Rows, chat Chat) {
	res := User{}
	response.Scan(&res.user_id, &res.user_name, &res.first_name)
	response.Close()

	err_count := 0
	if res.user_name != chat.user_name || res.first_name != chat.first_name {
		for {
			_, err := db.Exec("UPDATE shop_users SET user_name=$1, first_name=$2 WHERE user_id=$3", chat.user_name, chat.first_name, chat.user_id)
			if err == nil {
				break
			} else {
				err_count += 1
			}
			if err_count == 10 {
				break
			}
		}
	}
}

func SendOrdersStatus(response *sql.Rows, ChatID int64) {
	result := "[Статусы]"
	for response.Next() {
		res := Order{}
		response.Scan(&res.order_id, &res.status)
		result += fmt.Sprintf("\nЗаказ №%d - %s", res.order_id, status[res.status-1])
	}
	response.Close()

	if result == "[Статусы]" {
		go SendMessage(ChatID, "Нет активных заказов.")
	} else {
		go SendMessage(ChatID, result)
	}
}
