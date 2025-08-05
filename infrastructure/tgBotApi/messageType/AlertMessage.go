package messageType

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AlertMessage struct {
	Text       string
	Bot        tgbotapi.BotAPI
	CallBackID string
}

func (t AlertMessage) Send() tgbotapi.Message {
	msg := tgbotapi.NewCallbackWithAlert(t.CallBackID, t.Text)
	res, _ := t.Bot.Send(msg)

	return res
}
