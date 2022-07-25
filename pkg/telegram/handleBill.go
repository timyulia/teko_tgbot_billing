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
		if err != nil {
			return err
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "enter the amount")
		_, err = b.bot.Send(msg)
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
		_, err = b.bot.Send(msg)
		return err
	}
	err = b.repos.AddAmount(id, message.Chat.ID)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "error")
		_, err = b.bot.Send(msg)
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "describe")
	err = b.repos.UpdateCond(message.Chat.ID, "description")
	if err != nil {
		return err
	}
	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) handleDescription(message *tgbotapi.Message, sub string) error {

	err := b.repos.AddDescription(message.Text, message.Chat.ID, sub)
	if err != nil {
		return err
	}
	if sub == "description" {
		err = b.repos.UpdateCond(message.Chat.ID, "email")
		if err != nil {
			return err
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "enter email")
		_, err = b.bot.Send(msg)
	} else {
		err = b.repos.UpdateCond(message.Chat.ID, "validated")
		if err != nil {
			return err
		}
		err = b.buttonsSecond(message.Chat.ID, "created")
	}
	return err
}

func (b *Bot) handleHistory(message *tgbotapi.Message) error {
	bills, err := b.repos.GetHistory(message.Chat.ID)
	if err != nil {
		return err
	}
	if bills == nil {
		err = b.buttonsSecond(message.Chat.ID, "no bills yet")
		return err
	}
	res, _ := json.MarshalIndent(bills, "", "")
	err = b.buttonsSecond(message.Chat.ID, string(res))
	return err
}

func (b *Bot) handleTotal(message *tgbotapi.Message) error {
	res, err := b.repos.GetTotal(message.Chat.ID)
	if err != nil {
		return err
	}
	err = b.buttonsSecond(message.Chat.ID, fmt.Sprintf("%d", res))
	return err
}
