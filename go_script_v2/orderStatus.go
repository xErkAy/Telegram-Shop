package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

var status = [3]string{"не взят в работу", "готовится", "готов к выдаче"}

func SendOrdersStatus(ChatID int64) {
	var (
		orderId     int
		orderStatus int
	)
	response, _ := db.Query("SELECT order_id, status FROM shop_order WHERE user_id_id=$1 and is_closed=FALSE", ChatID)

	result := "[Статусы]"
	for response.Next() {
		response.Scan(&orderId, &orderStatus)
		result += fmt.Sprintf("\nЗаказ №%d - %s", orderId, status[orderStatus-1])
	}
	response.Close()

	if result == "[Статусы]" {
		go bot.SendMessage(ChatID, "Нет активных заказов.")
	} else {
		go bot.SendMessage(ChatID, result)
	}
}
