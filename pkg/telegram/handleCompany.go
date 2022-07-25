package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (b *Bot) handleLogIn(message *tgbotapi.Message) error {

	id, err := strconv.Atoi(message.Text)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "you have to enter a number")
		_, err = b.bot.Send(msg)
		return err
	}
	exists, err := b.repos.ValidateCompany(id, message.Chat.ID)
	if err != nil {
		return err
	}
	if exists {
		err = b.buttonsSecond(message.Chat.ID, "logged in")
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "there is no such company, try another one")
		_, err = b.bot.Send(msg)
	}
	return err
}

func (b *Bot) handleStartedK(message *tgbotapi.Message) error {
	var err error
	switch message.Text {
	case "add":
		err = b.repos.UpdateCond(message.Chat.ID, "adding")
		if err != nil {
			return err
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "enter the ID, it's a number")
		_, err = b.bot.Send(msg)
	case "log in":
		err = b.repos.UpdateCond(message.Chat.ID, "logging in")
		if err != nil {
			return err
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "enter the ID, it's a number")
		_, err = b.bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "you have to add or log in")
		_, err = b.bot.Send(msg)
	}
	return err
}

func (b *Bot) handleAddIdK(message *tgbotapi.Message) error {

	id, err := strconv.Atoi(message.Text)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "you have to enter a number")
		_, err = b.bot.Send(msg)
		return err
	}
	_, err = b.repos.AddCompany(id)
	if err != nil {
		return err
	}
	err = b.repos.UpdateCond(message.Chat.ID, "started")
	if err != nil {
		return err
	}
	err = b.buttonsFirst(message.Chat.ID, "added")
	return err
}
