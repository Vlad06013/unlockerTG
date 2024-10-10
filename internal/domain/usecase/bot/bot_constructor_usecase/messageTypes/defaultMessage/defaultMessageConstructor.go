package defaultMessage

import (
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DefaultMessageConstructor struct {
	Message    *entity.Message
	ChatId     int64
	Text       string
	Type       string
	Keyboard   entity.Keyboard
	Buttons    []tgbotapi.InlineKeyboardButton
	CallBackID *string
}

func (m *DefaultMessageConstructor) TypeMessage() string {
	return m.Type
}
func (m *DefaultMessageConstructor) GetMessage() *entity.Message {
	return m.Message
}

func (m *DefaultMessageConstructor) TextMessage() *string {
	return &m.Text
}

func (m *DefaultMessageConstructor) Chat() int64 {
	return m.ChatId
}

func (m *DefaultMessageConstructor) ButtonsMessage() []tgbotapi.InlineKeyboardButton {
	return m.Buttons
}

func (m *DefaultMessageConstructor) LastMessage() *entity.Message {
	return m.Message
}

func generateButtons(keyboard entity.Keyboard) []tgbotapi.InlineKeyboardButton {
	var buttons, convertedButtons []tgbotapi.InlineKeyboardButton
	prefix := "mess"

	if len(keyboard.Buttons) != 0 {
		for _, b := range keyboard.Buttons {
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(b.Text, prefix+"_"+b.CallbackData))
		}
		convertedButtons = tgbotapi.NewInlineKeyboardRow(buttons...)
	}
	return convertedButtons
}

func NewMessage(construct DefaultMessageConstructor) entity.Constructable {

	buttons := generateButtons(construct.Keyboard)

	var defaultMessage entity.Constructable = &DefaultMessageConstructor{
		Message:    construct.Message,
		ChatId:     construct.ChatId,
		Text:       construct.Text,
		Type:       construct.Type,
		Keyboard:   construct.Keyboard,
		Buttons:    buttons,
		CallBackID: nil,
	}

	return defaultMessage
}
