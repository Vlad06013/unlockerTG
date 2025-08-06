package messageType

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageWithImagesGroup struct {
	Text   string
	Bot    tgbotapi.BotAPI
	ChatId int64
	Media  []interface{}
}

func (t MessageWithImagesGroup) Send() tgbotapi.Message {

	msg := tgbotapi.NewMediaGroup(t.ChatId, t.Media)

	res, err := t.Bot.Send(msg)
	if err != nil {
		fmt.Println("sendError", err)
	}
	return res
}
