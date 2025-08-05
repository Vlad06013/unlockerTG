package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type NotFoundScreen struct {
	OutputMessage messageType.OutputMessage
}

func NotFound(user TgUser.TgUser, bot tgbotapi.BotAPI) BaseScreen {
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

func (c NotFoundScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c NotFoundScreen) GetScreenName() string {
	return "not_found"
}
