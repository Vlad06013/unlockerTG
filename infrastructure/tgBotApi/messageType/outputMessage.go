package messageType

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type OutputMessage interface {
	Send() tgbotapi.Message
}
