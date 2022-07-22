package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "there is no such command here")
	switch message.Command() {
	case "start":
		b.handleCommandStart(message)
		return nil
	case "reset":
		b.repos.UpdateCond(message.Chat.ID, "started")
		b.buttonsFirst(message.Chat.ID, "you are in the beginning now")
		return nil
	case "help":
		msg.Text = "you should log in existing company, then you will be able to create bills, see total amounts and " +
			"the history of bills\nalso you can add a new company"
		_, err := b.bot.Send(msg)
		return err
	case "back":
		b.handleCommandBack(message)
		return nil
	default:
		_, err := b.bot.Send(msg)
		return err
	}

}

func (b *Bot) handleCommandBack(message *tgbotapi.Message) {
	cond := b.repos.CheckCond(message.Chat.ID)
	switch cond {
	case "adding", "logging in", "validated":
		b.repos.UpdateCond(message.Chat.ID, "started")
		b.buttonsFirst(message.Chat.ID, "you're back to add a company or log in")

	case "amount", "description", "email":
		b.repos.UpdateCond(message.Chat.ID, "validated")
		b.buttonsSecond(message.Chat.ID, "you're back to choose an operation")

	default:
		b.buttonsFirst(message.Chat.ID, "you are in the beginning already")
	}

}
func (b *Bot) handleCommandStart(message *tgbotapi.Message) {
	b.repos.AddUser(message.Chat.ID)
	b.buttonsFirst(message.Chat.ID, "you can add a company and log in")
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	cond := b.repos.CheckCond(message.Chat.ID)
	switch cond {
	case "started":
		b.handleStartedK(message)
	case "adding":
		b.handleAddIdK(message)
	case "logging in":
		b.handleLogIn(message)
	case "validated":
		b.handleLogged(message)
	case "amount":
		b.handleAmount(message)
	case "description":
		b.handleDescription(message, "description")
	case "email":
		b.handleDescription(message, "email")
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "don't understand")
		b.bot.Send(msg)
	}
}
