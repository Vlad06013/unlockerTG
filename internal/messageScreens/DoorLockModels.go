package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/DoorLockModel"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

type DoorLockModelsScreen struct {
	OutputMessage messageType.OutputMessage
}

func DoorLockModels(user TgUser.TgUser, bot tgbotapi.BotAPI, db *gorm.DB, filter *uint64) BaseScreen {
	s := DoorLockModel.Storage{DB: db}

	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	backBtnCB := "doorLockMarks"
	doorsLockModels := s.GetByMarkId(*filter)
	text := "Выберете модель замка"

	if len(doorsLockModels) == 0 {
		text = "Модели отсутствуют"
	}
	rows := make([][]tgbotapi.InlineKeyboardButton, len(doorsLockModels)+1)

	for i := 0; i < len(doorsLockModels); i++ {
		callbackData := "doorLockModelsDetail|" + strconv.FormatUint(uint64(doorsLockModels[i].ID), 10)
		btnText := doorsLockModels[i].Name

		rows[i] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
			Text:         btnText,
			CallbackData: &callbackData,
		})
	}
	rows[len(doorsLockModels)] = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.InlineKeyboardButton{
			Text:         "Назад",
			CallbackData: &backBtnCB,
		},
	)

	buttons = rows

	keyboard = tgbotapi.NewInlineKeyboardMarkup(buttons...)

	var message = messageType.TextWithButtonsMessage{
		Text:    text,
		Bot:     bot,
		ChatId:  user.TgUserId,
		Buttons: keyboard,
	}

	var screen = DoorLockMarksScreen{
		OutputMessage: messageType.OutputMessage(message),
	}

	var baseScreen BaseScreen = screen

	return baseScreen
}

func (c DoorLockModelsScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c DoorLockModelsScreen) GetScreenName() string {
	return "doorLockModels"
}
