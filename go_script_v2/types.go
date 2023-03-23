package main

type Chat struct {
	ID        int64
	messageID int
	userId    string
	userName  string
	firstName string
}

type Notify struct {
	UserId  int64  `json:"user_id"`
	Message string `json:"message_text"`
}

type Response struct {
	Message string `json:"message"`
}
