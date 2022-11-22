package main

import (
	"HelpWithComm/pkg/handler"
	"HelpWithComm/pkg/repository"
	"HelpWithComm/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

/*const (
	CostGaz = 7.04
	CostLight = 4.36
	CostWater = 26.90
	Other = domPhone + Remont + JKX
	domPhone = 150
	Remont = 260
	JKX = 612
	OtherForInter = 1267
)*/

const Token = ""

func main() {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil{
		log.Panic(err)
	}

	bot.Debug = true

	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Panic(err)
	}

	rep := repository.NewRepository(db)
	handle := handler.NewNum(rep)
	TBot := telegram.NewBot(bot, handle)

	go func() {
		if err = TBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()
}
