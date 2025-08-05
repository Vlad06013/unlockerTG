package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarModel"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

type CarModelsScreen struct {
	OutputMessage messageType.OutputMessage
}

func CarModels(user TgUser.TgUser, bot tgbotapi.BotAPI, db *gorm.DB, filter *uint64) BaseScreen {
	s := CarModel.Storage{DB: db}

	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	backBtnCB := "carMarks"
	carModels := s.GetByMarkId(*filter)
	text := "Выберете модель авто"

	if len(carModels) == 0 {
		text = "Модели отсутствуют"
	}
	rows := make([][]tgbotapi.InlineKeyboardButton, len(carModels)+1)

	for i := 0; i < len(carModels); i++ {
		callbackData := "carModelDetail|" + strconv.FormatUint(uint64(carModels[i].ID), 10)
		btnText := carModels[i].Name

		rows[i] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
			Text:         btnText,
			CallbackData: &callbackData,
		})
	}
	rows[len(carModels)] = tgbotapi.NewInlineKeyboardRow(
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

func (c CarModelsScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c CarModelsScreen) GetScreenName() string {
	return "carModels"
}
