package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	Action *sql.DB
}

var database Database

func initDB() (Error error) {
	database.Action, Error = sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=shop_database sslmode=disable")
	return Error
}

func (database Database) isUserNew(chat Chat) (bool, error) {
	response, _ := database.Action.Query("SELECT * FROM shop_user WHERE user_id=$1", chat.userId)
	if !response.Next() {
		_, err := database.Action.Exec("INSERT INTO shop_user VALUES ($1, $2, $3, FALSE)", chat.userId, chat.userName, chat.firstName)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	database.updateUserInfo(response, chat)
	return false, nil
}

func (database Database) updateUserInfo(response *sql.Rows, chat Chat) {
	var (
		userName  string
		firstName string
	)
	response.Scan(&userName, &firstName)
	response.Close()

	if userName != chat.userName || firstName != chat.firstName {
		for i := 0; i < 10; i++ {
			_, err := database.Action.Exec("UPDATE shop_user SET user_name=$1, first_name=$2 WHERE user_id=$3", chat.userName, chat.firstName, chat.userId)
			if err == nil {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (database Database) isActiveChat(userId int64) (bool, int) {
	var (
		isChatActive bool
		orderId      int
	)

	response, _ := database.Action.Query("SELECT is_chat_active, order_id FROM shop_order WHERE user_id_id=$1 AND is_chat_active=TRUE", userId)
	response.Next()
	response.Scan(&isChatActive, &orderId)
	response.Close()

	return isChatActive, orderId
}

func (database Database) isOrderActive(userId int64) bool {
	var isOrderActive bool

	response, _ := database.Action.Query("SELECT is_order_active FROM shop_user WHERE user_id=$1", userId)
	response.Next()
	response.Scan(&isOrderActive)
	response.Close()

	return isOrderActive
}

func (database Database) makeOrderActive(userId int64) {
	_, err := database.Action.Exec("UPDATE shop_user SET is_order_active=TRUE WHERE user_id=$1", userId)
	if err != nil {
		go bot.SendMessage(userId, "Произошла ошибка. Попробуйте еще раз.")
	} else {
		go bot.SendMessage(userId, "Пожалуйста, введите Ваш заказ ОДНИМ сообщением.")
	}
}

func (database Database) makeNewOrder(userId int64, orderValue string) {
	_, err := database.Action.Exec("UPDATE shop_user SET is_order_active=FALSE WHERE user_id=$1", userId)
	if err != nil {
		go bot.SendMessage(userId, "Произошла ошибка. Попробуйте еще раз.")
		return
	}

	formData := url.Values{
		"user_id":     {strconv.FormatInt(userId, 10)},
		"order_value": {orderValue},
	}
	response, err := http.PostForm("http://localhost:8000/api/orders/", formData)

	if err != nil {
		go bot.SendMessage(userId, "Произошла ошибка при создании заказа. Попробуйте еще раз.")
		return
	}

	var message Response
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &message)

	go bot.SendMessage(userId, message.Message)
}
