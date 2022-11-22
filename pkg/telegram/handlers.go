package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (b *Bot) handlerCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		return b.handlerStartCommand(message)
	case "help":
		return b.handlerHelpCommand(message)
	default:
		return b.handlerUnknownCommand(message)
	}
}

func (b *Bot) handlerMessage(message *tgbotapi.Message) error {
	if err := b.ValidMessage(message); err != nil {
		return err
	}
//	id, err := b.GetId(message)
//	if err != nil {
//		return err
//	}

	id := 1
	text := "Сумма: "

	sum, err := b.handler.Count(id, message.Text)
	if err != nil {
		return err
	}
	text += strconv.Itoa(sum)

	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handlerStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Приветствую, для начала работы введите /help")

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handlerHelpCommand(message *tgbotapi.Message) error{
	msg := tgbotapi.NewMessage(message.Chat.ID, "Введите показания счетчиков, затем выберете номер квартиры.")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handlerUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Такая команда не зарегистрирована")

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) ValidMessage(message *tgbotapi.Message) error {
	return nil
}

//func (b *Bot) GetId(message *tgbotapi.Message) (int, error) {
//	return 0, nil
//}
