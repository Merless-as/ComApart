package telegram

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error{
	return nil
}

func (b *Bot) handlerUpdates(updates tgbotapi.UpdatesChannel) {

}

func (b *Bot) initUpdateChannel() error {
	return nil
}