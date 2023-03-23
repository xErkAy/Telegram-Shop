package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Bot struct {
	Action *tgbotapi.BotAPI
}

var bot Bot

var keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сделать заказ", "makeorder"),
		tgbotapi.NewInlineKeyboardButtonData("Посмотреть меню", "getmenu"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Посмотреть статусы заказов", "getordersstatus"),
	),
)

func initBot() error {
	var err error

	bot.Action, err = tgbotapi.NewBotAPI("5288667624:AAHkmEmA-SYeLHo1RS3TC1gpHuh9SlI6V3g")
	if err != nil {
		return err
	}
	bot.Action.Debug = true
	return nil
}

func (bot Bot) SendKeyboard(ChatID int64) {
	message := tgbotapi.NewMessage(ChatID, "Что вы хотите сделать?")
	message.ReplyMarkup = keyboard
	bot.Action.Send(message)
}

func (bot Bot) SendMessage(ChatID int64, message string) {
	bot.Action.Send(tgbotapi.NewMessage(ChatID, message))
}

func (bot Bot) ReplyToMessageID(ChatID int64, messageID int, Text string) {
	message := tgbotapi.NewMessage(ChatID, Text)
	message.ReplyToMessageID = messageID
	bot.Action.Send(message)
}

func (bot Bot) SendPhoto(ChatID int64, FilePath string) {
	bot.Action.Send(tgbotapi.NewPhoto(ChatID, tgbotapi.FilePath(FilePath)))
}

func (bot Bot) ReplyWithPhoto(ChatID int64, messageID int, FilePath string) {
	message := tgbotapi.NewPhoto(ChatID, tgbotapi.FilePath(FilePath))
	message.ReplyToMessageID = messageID
	bot.Action.Send(message)
}

func (bot Bot) SendDocument(ChatID int64, FilePath string) {
	bot.Action.Send(tgbotapi.NewDocument(ChatID, tgbotapi.FilePath(FilePath)))
}

func (bot Bot) ReplyWithDocument(ChatID int64, messageID int, FilePath string) {
	message := tgbotapi.NewDocument(ChatID, tgbotapi.FilePath(FilePath))
	message.ReplyToMessageID = messageID
	bot.Action.Send(message)
}

func (bot Bot) SendSticker(ChatID int64, FilePath string) {
	bot.Action.Send(tgbotapi.NewSticker(ChatID, tgbotapi.FilePath(FilePath)))
}
