package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) buttonsSecond(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	reply := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("create a bill"),
			tgbotapi.NewKeyboardButton("last 10 bills"),
			tgbotapi.NewKeyboardButton("total amount"),
		),
	)
	reply.OneTimeKeyboard = true
	msg.ReplyMarkup = reply
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) buttonsFirst(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	reply := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("add"),
			tgbotapi.NewKeyboardButton("log in"),
		),
	)
	reply.OneTimeKeyboard = true
	msg.ReplyMarkup = reply
	_, err := b.bot.Send(msg)
	return err
}
