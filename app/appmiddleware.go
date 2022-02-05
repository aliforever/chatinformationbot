package app

import tgbotapi "github.com/aliforever/go-telegram-bot-api"

func (app App) Middleware(update *tgbotapi.Update) (ignoreUpdate bool) {
	return
}
