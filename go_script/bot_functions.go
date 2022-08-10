package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SendHelpMessage(ChatID int64) {
	bot.Send(tgbotapi.NewMessage(ChatID, "Доступные команды:\n/help - справка\n/menu - получить меню\n/start - оформить заказ"))
}

func SendMessage(ChatID int64, message string) {
	bot.Send(tgbotapi.NewMessage(ChatID, message))
}

func ReplyToMessageID(ChatID int64, messageID int, Text string) {
	message := tgbotapi.NewMessage(ChatID, Text)
	message.ReplyToMessageID = messageID
	bot.Send(message)
}

func SendPhoto(ChatID int64, FilePath string) {
	bot.Send(tgbotapi.NewPhoto(ChatID, tgbotapi.FilePath(FilePath)))
}

func ReplyWithPhoto(ChatID int64, messageID int, FilePath string) {
	message := tgbotapi.NewPhoto(ChatID, tgbotapi.FilePath(FilePath))
	message.ReplyToMessageID = messageID
	bot.Send(message)
}

func SendDocument(ChatID int64, FilePath string) {
	bot.Send(tgbotapi.NewDocument(ChatID, tgbotapi.FilePath(FilePath)))
}

func ReplyWithDocument(ChatID int64, messageID int, FilePath string) {
	message := tgbotapi.NewDocument(ChatID, tgbotapi.FilePath(FilePath))
	message.ReplyToMessageID = messageID
	bot.Send(message)
}

func SendSticker(ChatID int64, FilePath string) {
	bot.Send(tgbotapi.NewSticker(ChatID, tgbotapi.FilePath(FilePath)))
}
