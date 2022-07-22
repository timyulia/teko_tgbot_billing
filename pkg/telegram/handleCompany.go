package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (b *Bot) handleLogIn(message *tgbotapi.Message) error {

	id, err := strconv.Atoi(message.Text)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "you have to enter a number")
		b.bot.Send(msg)
		return err
	}
	exists, err := b.repos.ValidateCompany(id, message.Chat.ID)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "error")
		_, err = b.bot.Send(msg)
		return err
	}
	if exists {
		b.buttonsSecond(message.Chat.ID, "logged in")
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "there is no such company, try another one")
		b.bot.Send(msg)
	}

	return nil
}

func (b *Bot) handleStartedK(message *tgbotapi.Message) {

	switch message.Text {
	case "add":
		b.repos.UpdateCond(message.Chat.ID, "adding")
		msg := tgbotapi.NewMessage(message.Chat.ID, "enter the ID, it's a number")
		b.bot.Send(msg)
	case "log in":
		b.repos.UpdateCond(message.Chat.ID, "logging in")
		msg := tgbotapi.NewMessage(message.Chat.ID, "enter the ID, it's a number")
		b.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "you have to add or log in")
		b.bot.Send(msg)
	}
}

func (b *Bot) handleAddIdK(message *tgbotapi.Message) error {

	id, err := strconv.Atoi(message.Text)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "you have to enter a number")
		b.bot.Send(msg)
		return err
	}
	_, err = b.repos.AddCompany(id)
	if err != nil {
		return err
	}
	b.repos.UpdateCond(message.Chat.ID, "started")
	msg := tgbotapi.NewMessage(message.Chat.ID, "added")
	reply := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("add"),
			tgbotapi.NewKeyboardButton("log in"),
		),
	)
	reply.OneTimeKeyboard = true
	msg.ReplyMarkup = reply
	_, err = b.bot.Send(msg)
	return err
}
