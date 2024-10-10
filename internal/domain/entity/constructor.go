package entity

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BotApi struct {
	Api tgbotapi.BotAPI
	Bot *Bot
}

type Constructable interface {
	TypeMessage() string
	TextMessage() *string
	Chat() int64
	ButtonsMessage() []tgbotapi.InlineKeyboardButton
	GetMessage() *Message
}
