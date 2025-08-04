package messageType

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TextWithButtonsMessage struct {
	Text    string
	Bot     tgbotapi.BotAPI
	ChatId  int64
	Buttons tgbotapi.InlineKeyboardMarkup
	Screen  string
}

func (t TextWithButtonsMessage) Send() tgbotapi.Message {

	msg := tgbotapi.NewMessage(t.ChatId, t.Text)
	msg.ParseMode = "HTML"
	buttons := t.Buttons

	if len(buttons.InlineKeyboard) != 0 {
		msg.ReplyMarkup = buttons
	}
	res, err := t.Bot.Send(msg)
	if err != nil {
		fmt.Println("sendError", err)
	}
	return res
}
