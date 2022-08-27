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

var db, db_err = sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=shop_database sslmode=disable")

func isUserNew(chat Chat) (bool, error) {
	response, _ := db.Query("SELECT * FROM shop_users WHERE user_id=$1", chat.user_id)
	if !response.Next() {
		_, err := db.Exec("INSERT INTO shop_users VALUES ($1, $2, $3, FALSE)", chat.user_id, chat.user_name, chat.first_name)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	go updateUserInfo(response, chat)
	return false, nil
}

func updateUserInfo(response *sql.Rows, chat Chat) {
	var (
		user_name  string
		first_name string
	)
	response.Scan(&user_name, &first_name)
	response.Close()

	if user_name != chat.user_name || first_name != chat.first_name {
		for i := 0; i < 10; i++ {
			_, err := db.Exec("UPDATE shop_users SET user_name=$1, first_name=$2 WHERE user_id=$3", chat.user_name, chat.first_name, chat.user_id)
			if err == nil {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func isActiveChat(user_id int64) (bool, int) {
	var (
		is_chat_active bool
		order_id       int
	)

	response, _ := db.Query("SELECT is_chat_active, order_id FROM shop_orders WHERE user_id_id=$1 AND is_chat_active=TRUE", user_id)
	response.Next()
	response.Scan(&is_chat_active, &order_id)
	response.Close()

	return is_chat_active, order_id
}

func isOrderActive(user_id int64) bool {
	var is_order_active bool

	response, _ := db.Query("SELECT is_order_active FROM shop_users WHERE user_id=$1", user_id)
	response.Next()
	response.Scan(&is_order_active)
	response.Close()

	return is_order_active
}

func makeOrderActive(user_id int64) {
	_, err := db.Exec("UPDATE shop_users SET is_order_active=TRUE WHERE user_id=$1", user_id)
	if err != nil {
		go SendMessage(user_id, "Произошла ошибка. Попробуйте еще раз.")
	} else {
		go SendMessage(user_id, "Пожалуйста, введите Ваш заказ ОДНИМ сообщением.")
	}
}

func makeNewOrder(user_id int64, order_value string) {
	_, err := db.Exec("UPDATE shop_users SET is_order_active=FALSE WHERE user_id=$1", user_id)
	if err != nil {
		go SendMessage(user_id, "Произошла ошибка. Попробуйте еще раз.")
		return
	}

	formData := url.Values{
		"user_id":     {strconv.FormatInt(user_id, 10)},
		"order_value": {order_value},
	}
	response, err := http.PostForm("http://localhost:8000/api/orders/", formData)

	if err != nil {
		go SendMessage(user_id, "Произошла ошибка при создании заказа. Попробуйте еще раз.")
		return
	}

	var message Response
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &message)

	go SendMessage(user_id, message.Message)
}
