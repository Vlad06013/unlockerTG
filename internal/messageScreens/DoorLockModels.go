package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/DoorLockModel"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type DoorLockModelsScreen struct {
	OutputMessage messageType.OutputMessage
}

func DoorLockModels(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	s := DoorLockModel.Storage{DB: dto.DB}

	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	backBtnCB := "doorLockMarks"
	doorsLockModels := s.GetByMarkId(1)
	text := "Выберете модель замка"

	if len(doorsLockModels) == 0 {
		text = "Модели отсутствуют"
	}
	rows := make([][]tgbotapi.InlineKeyboardButton, len(doorsLockModels)+1)

	for i := 0; i < len(doorsLockModels); i++ {
		callbackData := "doorLockModelsDetail|id_" + strconv.FormatUint(uint64(doorsLockModels[i].ID), 10)
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
		Bot:     dto.Bot,
		ChatId:  dto.User.TgUserId,
		Buttons: keyboard,
	}

	var screen = DoorLockMarksScreen{
		OutputMessage: messageType.OutputMessage(message),
	}

	var baseScreenInterface BaseScreen = screen

	return baseScreenInterface, true
}

func (c DoorLockModelsScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c DoorLockModelsScreen) GetScreenName() string {
	return "doorLockModels"
}
