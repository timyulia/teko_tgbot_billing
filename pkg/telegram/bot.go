package telegram

import (
	"accounting_teko/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI

	repos *repository.Repository
}

func NewBot(bot *tgbotapi.BotAPI, repos *repository.Repository) *Bot {
	return &Bot{bot: bot, repos: repos}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err := b.handleCommand(update.Message)
			if err != nil {
				b.handleError(update.Message, err)
			}
			continue
		}

		err := b.handleMessage(update.Message)
		if err != nil {
			b.handleError(update.Message, err)
		}
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleError(message *tgbotapi.Message, err error) {
	log.Print(err)
	msg := tgbotapi.NewMessage(message.Chat.ID, "something went wrong")
	b.bot.Send(msg)
}
