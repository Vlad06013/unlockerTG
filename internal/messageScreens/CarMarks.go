package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarMark"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

type CarMarksScreen struct {
	OutputMessage messageType.OutputMessage
}

func CarMarks(user TgUser.TgUser, bot tgbotapi.BotAPI, db *gorm.DB) BaseScreen {
	s := CarMark.Storage{DB: db}

	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	backBtnCB := "categories"
	carMarks := s.GetAll()

	//if len(doorsLockMarks) == 0 {
	//	text = "Нет подключенных доменов"
	//}
	rows := make([][]tgbotapi.InlineKeyboardButton, len(carMarks)+1)

	for i := 0; i < len(carMarks); i++ {
		callbackData := "carModels|" + strconv.FormatUint(uint64(carMarks[i].ID), 10)
		btnText := carMarks[i].Name

		rows[i] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
			Text:         btnText,
			CallbackData: &callbackData,
		})
	}
	rows[len(carMarks)] = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.InlineKeyboardButton{
			Text:         "Назад",
			CallbackData: &backBtnCB,
		},
	)

	buttons = rows

	keyboard = tgbotapi.NewInlineKeyboardMarkup(buttons...)

	var message = messageType.TextWithButtonsMessage{
		Text:    "Выберете марку",
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

func (c CarMarksScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c CarMarksScreen) GetScreenName() string {
	return "carMarks"
}
