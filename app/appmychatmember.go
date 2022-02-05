package app

import (
	"fmt"

	tgbotapi "github.com/aliforever/go-telegram-bot-api"
)

func (app App) MyChatMemberHandler(update *tgbotapi.Update) {
	if update.MyChatMember != nil && update.MyChatMember.NewChatMember != nil && update.MyChatMember.Chat != nil {
		chat := update.MyChatMember.Chat
		text := "📜 Chat Information:\n\n"
		text += fmt.Sprintf("🔑 ID: %d\n", chat.Id)
		if chat.Title != "" {
			text += fmt.Sprintf("💎 Title: %s\n", chat.Title)
		}
		text += fmt.Sprintf("🛡 Type: %s\n", chat.Type)
		if chat.Username != "" {
			text += fmt.Sprintf("\U0001FAA7 Username: %s", chat.Username)
		}
		app.Send(app.Message().SetChatId(chat.Id).SetText(text))
		// TODO: Leave Chat
	}
}
