package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CategoryScreen struct {
	OutputMessage messageType.OutputMessage
}

func Categories(user TgUser.TgUser, bot tgbotapi.BotAPI) BaseScreen {
	cars := "carMarks"
	doors := "doorLockMarks"

	var keyboard tgbotapi.InlineKeyboardMarkup
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
		Bot:     bot,
		ChatId:  user.TgUserId,
		Buttons: keyboard,
	}

	var screen = CategoryScreen{
		OutputMessage: messageType.OutputMessage(message),
	}

	var baseScreen BaseScreen = screen

	return baseScreen
}

func (c CategoryScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c CategoryScreen) GetScreenName() string {
	return "categories"
}
