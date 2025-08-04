package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/DoorLockMark"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

type DoorLockMarksScreen struct {
	OutputMessage messageType.OutputMessage
}

func DoorLockMarks(user TgUser.TgUser, bot tgbotapi.BotAPI, db *gorm.DB) BaseScreen {
	s := DoorLockMark.Storage{DB: db}

	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	backBtnCB := "categories"
	doorsLockMarks := s.GetAll()

	//if len(doorsLockMarks) == 0 {
	//	text = "Нет подключенных доменов"
	//}
	rows := make([][]tgbotapi.InlineKeyboardButton, len(doorsLockMarks)+1)

	for i := 0; i < len(doorsLockMarks); i++ {
		callbackData := "doorLockModels|" + strconv.FormatUint(uint64(doorsLockMarks[i].ID), 10)
		btnText := doorsLockMarks[i].Name

		rows[i] = tgbotapi.NewInlineKeyboardRow(tgbotapi.InlineKeyboardButton{
			Text:         btnText,
			CallbackData: &callbackData,
		})
	}
	rows[len(doorsLockMarks)] = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.InlineKeyboardButton{
			Text:         "Назад",
			CallbackData: &backBtnCB,
		},
	)

	buttons = rows

	keyboard = tgbotapi.NewInlineKeyboardMarkup(buttons...)

	var message = messageType.TextWithButtonsMessage{
		Text:    "Выберете производителя замка",
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

func (c DoorLockMarksScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c DoorLockMarksScreen) GetScreenName() string {
	return "doorLockMarks"
}
