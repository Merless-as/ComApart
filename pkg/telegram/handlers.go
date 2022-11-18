package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) handlerCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		return b.handlerStartCommand(message)
	case "count":
		return b.handlerCountCommand(message)
	default:
		return b.handlerUnknownCommand(message)
	}
}

func (b *Bot) handlerMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handlerStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Приветствую, для начала работы введите /help")

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handlerCountCommand(message *tgbotapi.Message) error {
	return nil
}

func (b *Bot) handlerUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Такая команда не зарегистрирована")

	_, err := b.bot.Send(msg)
	return err
}
