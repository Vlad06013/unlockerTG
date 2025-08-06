package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CategoryScreen struct {
	OutputMessage messageType.OutputMessage
}

func NewCategoriesScreen(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	cars := "carMarks"
	doors := "doorLockMarks"

	rows := make([][]tgbotapi.InlineKeyboardButton, 1)

	rows[0] = tgbotapi.NewInlineKeyboardRow(

		tgbotapi.InlineKeyboardButton{
			Text:         "Дверные замки",
			CallbackData: &doors,
		},
		tgbotapi.InlineKeyboardButton{
			Text:         "Авто",
			CallbackData: &cars,
		},
	)

	keyboard = tgbotapi.NewInlineKeyboardMarkup(rows...)

	var message = messageType.TextWithButtonsMessage{
		Text:    "Выберете категорию",
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

func (c CategoryScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c CategoryScreen) GetScreenName() string {
	return "categories"
}
