package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NotFoundScreen struct {
	OutputMessage messageType.OutputMessage
}

func NewNotFound(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	cb := "categories"

	var keyboard tgbotapi.InlineKeyboardMarkup
	rows := make([][]tgbotapi.InlineKeyboardButton, 1)

	rows[0] = tgbotapi.NewInlineKeyboardRow(

		tgbotapi.InlineKeyboardButton{
			Text:         "В начало",
			CallbackData: &cb,
		},
	)

	keyboard = tgbotapi.NewInlineKeyboardMarkup(rows...)

	var message = messageType.TextWithButtonsMessage{
		Text:    "К сожалению ничего не найдено",
		Bot:     dto.Bot,
		ChatId:  dto.User.TgUserId,
		Buttons: keyboard,
	}

	var screen = CategoryScreen{
		OutputMessage: messageType.OutputMessage(message),
	}

	var baseScreenInterface BaseScreen = screen

	return baseScreenInterface, true
}

func (c NotFoundScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c NotFoundScreen) GetScreenName() string {
	return "not_found"
}
