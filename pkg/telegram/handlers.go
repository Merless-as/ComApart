package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) handlerCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		return b.handlerStartCommand(message)
	default:
		return b.handlerUnknownCommand(message)
	}
}

func (b *Bot) handlerMessage(message *tgbotapi.Message) error {
	return nil
}

func (b *Bot) handlerStartCommand(message *tgbotapi.Message) error {
	return nil
}

func (b *Bot) handlerUnknownCommand(message *tgbotapi.Message) error {
	return nil
}
