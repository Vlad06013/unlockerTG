package messageType

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TextMessage struct {
	Text   string
	Bot    tgbotapi.BotAPI
	ChatId int64
	Screen string
}

func (t TextMessage) Send() tgbotapi.Message {

	msg := tgbotapi.NewMessage(t.ChatId, t.Text)
	msg.ParseMode = "HTML"
	res, err := t.Bot.Send(msg)
	if err != nil {
		fmt.Println("sendError", err)
	}
	return res
}
