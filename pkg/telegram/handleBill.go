package telegram

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (b *Bot) handleLogged(message *tgbotapi.Message) error {
	var err error
	switch message.Text {
	case "create a bill":
		err = b.repos.UpdateCond(message.Chat.ID, "amount")
		msg := tgbotapi.NewMessage(message.Chat.ID, "enter the amount")
		b.bot.Send(msg)
	case "last 10 bills":
		err = b.handleHistory(message)
	case "total amount":
		err = b.handleTotal(message)
	default:
		err = b.buttonsSecond(message.Chat.ID, "you can create a bill, check last 10 bills, see a total amount")
	}
	return err
}

func (b *Bot) handleAmount(message *tgbotapi.Message) error {
	id, err := strconv.Atoi(message.Text)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "you have to enter a number")
		b.bot.Send(msg)
		return err
	}
	err = b.repos.AddAmount(id, message.Chat.ID)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "error")
		_, err = b.bot.Send(msg)
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "describe")
	b.bot.Send(msg)
	b.repos.UpdateCond(message.Chat.ID, "description")
	return nil
}

func (b *Bot) handleDescription(message *tgbotapi.Message, sub string) error {

	err := b.repos.AddDescription(message.Text, message.Chat.ID, sub)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "error")
		_, err = b.bot.Send(msg)
		return err
	}
	if sub == "description" {
		b.repos.UpdateCond(message.Chat.ID, "email")
		msg := tgbotapi.NewMessage(message.Chat.ID, "enter email")
		b.bot.Send(msg)
	} else {
		b.repos.UpdateCond(message.Chat.ID, "validated")
		b.buttonsSecond(message.Chat.ID, "created")
	}
	return nil
}

func (b *Bot) handleHistory(message *tgbotapi.Message) error {
	bills, err := b.repos.GetHistory(message.Chat.ID)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, err.Error())
		b.bot.Send(msg)
		return err
	}
	if bills == nil {
		b.buttonsSecond(message.Chat.ID, "no bills yet")
		return nil
	}
	res, _ := json.MarshalIndent(bills, "", "")
	b.buttonsSecond(message.Chat.ID, string(res))
	return nil
}

func (b *Bot) handleTotal(message *tgbotapi.Message) error {
	res, err := b.repos.GetTotal(message.Chat.ID)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, err.Error())
		b.bot.Send(msg)
		return err
	}
	b.buttonsSecond(message.Chat.ID, fmt.Sprintf("%d", res))
	return nil
}
