package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/DoorLockMark"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type DoorLockMarksScreen struct {
	OutputMessage messageType.OutputMessage
}

func NewDoorLockMarks(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	s := DoorLockMark.Storage{DB: dto.DB}

	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	backBtnCB := "categories"
	doorsLockMarks := s.GetAll()

	var text string
	if len(doorsLockMarks) == 0 {
		text = "В процессе заполнения. Обратитесь позже"
	} else {
		text = "Выберете производителя замка"
	}
	rows := make([][]tgbotapi.InlineKeyboardButton, len(doorsLockMarks)+1)

	for i := 0; i < len(doorsLockMarks); i++ {
		callbackData := "doorLockModels|id_" + strconv.FormatUint(uint64(doorsLockMarks[i].ID), 10)
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

func (c DoorLockMarksScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c DoorLockMarksScreen) GetScreenName() string {
	return "doorLockMarks"
}
