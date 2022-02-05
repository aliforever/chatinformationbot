package main

import (
	"fmt"

	"github.com/aliforever/chatinformationbot/app"
	"github.com/aliforever/go-telebot"
)

func main() {
	token := "" // Place Token Here

	var (
		bot *telebot.Bot
		err error
	)

	bot, _, err = telebot.NewBot(token, app.App{}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = bot.Poll()
	if err != nil {
		fmt.Println(err)
		return
	}
}
