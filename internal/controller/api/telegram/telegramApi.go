package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//type Constructable interface {
//	TypeMessage() string
//	TextMessage() *string
//	ButtonsMessage() []tgbotapi.InlineKeyboardButton
//	Chat() int64
//	LastMessage() *entity.Message
//}
//type Output struct {
//	Constructable
//	Bot tgbotapi.BotAPI
//}

type Output struct {
	Text        *string
	TypeMessage string
	Bot         tgbotapi.BotAPI
	Buttons     []tgbotapi.InlineKeyboardButton
	ChatId      int64
}

func (o *Output) sendTextMessage() *tgbotapi.Message {

	msg := tgbotapi.NewMessage(o.ChatId, *o.Text)
	msg.ParseMode = "HTML"
	buttons := o.Buttons
	if len(buttons) != 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)
	}

	res, err := o.Bot.Send(msg)
	if err != nil {
		fmt.Println("sendError", err)
	}

	return &res
}
func (o *Output) DeleteMessage(messageID int) *tgbotapi.Message {
	msg := tgbotapi.NewDeleteMessage(o.ChatId, messageID)
	o.Bot.Send(msg)
	return nil
}

func (o *Output) SendMessage() *tgbotapi.Message {
	var send *tgbotapi.Message
	if o.TypeMessage == "message" {
		send = o.sendTextMessage()
	}
	//if o.Type == "alert" {
	//	o.sendAlert(o.MessageConstructor.CallBackID, o.MessageConstructor.Text)
	//	sent = nil
	//}

	//res := o.sendAnimation(chatId)
	return send

}

//func NewOutput(o *Output) *Output {
//	var output = &Output{
//		Text:        o.Text,
//		TypeMessage: o.TypeMessage,
//		Bot: o.Bot,
//		Buttons: o.Buttons,
//		ChatId: o.ChatId,
//	}
//	return output
//}
