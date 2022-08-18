package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

var status = [3]string{"не взят в работу", "готовится", "готов к выдаче"}

func SendOrdersStatus(ChatID int64) {
	response, _ := db.Query("SELECT order_id, status FROM shop_orders WHERE user_id_id=$1 and is_closed=FALSE", ChatID)

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
