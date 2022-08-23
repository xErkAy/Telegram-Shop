package main

type Chat struct {
	ID         int64
	messageID  int
	user_id    string
	user_name  string
	first_name string
}

type User struct {
	user_id    string
	user_name  string
	first_name string
}

type Notify struct {
	User_id int64  `json:"user_id"`
	Message string `json:"message_text"`
}

type Order struct {
	user_id            int64
	order_id           int
	status             int
	is_chat_active     bool
	is_order_receiving bool
}
