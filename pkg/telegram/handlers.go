package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "there is no such command here")
	switch message.Command() {
	case "start":
		err := b.handleCommandStart(message)
		return err
	case "reset":
		err := b.repos.UpdateCond(message.Chat.ID, "started")
		if err != nil {
			return err
		}
		err = b.buttonsFirst(message.Chat.ID, "you are in the beginning now")
		return err
	case "help":
		msg.Text = "you should log in existing company, then you will be able to create bills, see total amounts and " +
			"the history of bills\nalso you can add a new company"
		_, err := b.bot.Send(msg)
		return err
	case "back":
		err := b.handleCommandBack(message)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}

}

func (b *Bot) handleCommandBack(message *tgbotapi.Message) error {
	cond, err := b.repos.CheckCond(message.Chat.ID)
	if err != nil {
		return err
	}
	switch cond {
	case "adding", "logging in", "validated":
		err = b.repos.UpdateCond(message.Chat.ID, "started")
		if err != nil {
			return err
		}
		err = b.buttonsFirst(message.Chat.ID, "you're back to add a company or log in")

	case "amount", "description", "email":
		err = b.repos.UpdateCond(message.Chat.ID, "validated")
		if err != nil {
			return err
		}
		err = b.buttonsSecond(message.Chat.ID, "you're back to choose an operation")

	default:
		err = b.buttonsFirst(message.Chat.ID, "you are in the beginning already")
	}
	return err
}

func (b *Bot) handleCommandStart(message *tgbotapi.Message) error {
	err := b.repos.AddUser(message.Chat.ID)
	if err != nil {
		return err
	}
	err = b.buttonsFirst(message.Chat.ID, "you can add a company and log in")
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	cond, err := b.repos.CheckCond(message.Chat.ID)
	if err != nil {
		return err
	}
	switch cond {
	case "started":
		err = b.handleStartedK(message)
	case "adding":
		err = b.handleAddIdK(message)
	case "logging in":
		err = b.handleLogIn(message)
	case "validated":
		err = b.handleLogged(message)
	case "amount":
		err = b.handleAmount(message)
	case "description":
		err = b.handleDescription(message, "description")
	case "email":
		err = b.handleDescription(message, "email")
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "don't understand")
		_, err = b.bot.Send(msg)
	}
	return err
}
