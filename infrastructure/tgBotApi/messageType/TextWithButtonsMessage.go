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
}

func (t TextWithButtonsMessage) Send() tgbotapi.Message {

	msg := tgbotapi.NewMessage(t.ChatId, t.Text)
	msg.ParseMode = "HTML"

	if len(t.Buttons.InlineKeyboard) != 0 {
		msg.ReplyMarkup = t.Buttons
	}
	res, err := t.Bot.Send(msg)
	if err != nil {
		fmt.Println("sendError", err)
	}
	return res
}
