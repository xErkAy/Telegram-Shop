package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	_ "github.com/lib/pq"
)

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

	SendMessage(int64(chat.User_id), chat.Message)
	connection.Close()
}
