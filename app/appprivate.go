package app

import (
	"fmt"

	tgbotapi "github.com/aliforever/go-telegram-bot-api"
)

func (app App) Welcome(update *tgbotapi.Update, isSwitched bool) (newState string) {
	if update.Message != nil && update.Message.Text != "" {
		if !isSwitched {
			if update.Message.Text == "Hello" {
				newState = "Hello"
				return
			}
		}
	}

	keyboard := app.Tools.Keyboards.NewReplyKeyboardFromSlicesOfStrings([][]string{
		{"Hello"},
	}).SetResizeKeyboard(true)
	app.Send(app.Message().SetText(fmt.Sprintf("Hello World!\nYour name is: %s", update.Message.Chat.FirstName)).SetReplyMarkup(keyboard))
	return
}

func (app App) Hello(update *tgbotapi.Update, isSwitched bool) (newState string) {
	if !isSwitched {
		if update.Message.Text == "Back" {
			newState = "Welcome"
			return
		}
	}
	keyboard := app.Tools.Keyboards.NewReplyKeyboardFromSlicesOfStrings([][]string{
		{"Back"},
	}).SetResizeKeyboard(true)
	app.Send(app.Message().SetText(fmt.Sprintf("You are in Hello Menu %s", update.Message.Chat.FirstName)).SetReplyMarkup(keyboard))
	return
}
