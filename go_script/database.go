package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var db, db_err = sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=shop_database sslmode=disable")

func isUserNew(response *sql.Rows, chat Chat) bool {
	if !response.Next() {
		for i := 0; i < 10; i++ {
			_, err := db.Exec("INSERT INTO shop_users VALUES ($1, $2, $3)", chat.user_id, chat.user_name, chat.first_name)
			if err == nil {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		return true
	}
	return false
}

func updateUserInfo(response *sql.Rows, chat Chat) {
	res := User{}
	response.Scan(&res.user_id, &res.user_name, &res.first_name)
	response.Close()

	if res.user_name != chat.user_name || res.first_name != chat.first_name {
		for i := 0; i < 10; i++ {
			_, err := db.Exec("UPDATE shop_users SET user_name=$1, first_name=$2 WHERE user_id=$3", chat.user_name, chat.first_name, chat.user_id)
			if err == nil {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func isActiveChat(chat Chat) (bool, int) {
	res := Order{}

	response, _ := db.Query("SELECT is_chat_active, order_id FROM shop_orders WHERE user_id_id=$1 AND is_chat_active=TRUE", chat.user_id)
	response.Next()
	response.Scan(&res.is_chat_active, &res.order_id)
	response.Close()

	return res.is_chat_active, res.order_id
}
