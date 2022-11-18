package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	log.Println("Bot started")

	updates, err := b.initUpdateChannel()
	if err != nil {
		return err
	}

	b.handlerUpdates(updates)
	return nil
}

func (b *Bot) handlerUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handlerCommand(update.Message); err != nil {
				continue
			}
		}
		if err := b.handlerMessage(update.Message); err != nil {
			continue
		}
	}
}

func (b *Bot) initUpdateChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
