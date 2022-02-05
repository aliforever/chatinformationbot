package app

import (
	tgbotapi "github.com/aliforever/go-telegram-bot-api"
)

func (app App) ChatMemberHandler(update *tgbotapi.Update) {
	// j, _ := json.Marshal(update)
	// fmt.Println("chat_member update", string(j))
}
