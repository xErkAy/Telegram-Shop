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
	User_id int    `json:"user_id"`
	Message string `json:"message_text"`
}
