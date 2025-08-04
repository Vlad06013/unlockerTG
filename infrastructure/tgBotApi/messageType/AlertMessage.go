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
	//if err != nil {
	//	fmt.Println("sendError", err)
	//}
	return res
}

//func (t AlertMessage) GetScreenName() string {
//	return t.Screen
//}

func (t AlertMessage) MessageType() OutputMessage {
	var output OutputMessage = t
	return output
}
