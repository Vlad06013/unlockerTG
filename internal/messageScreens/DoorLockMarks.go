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

	page, _ := strconv.ParseUint(dto.Filter["page"], 10, 32)

	var row []tgbotapi.InlineKeyboardButton
	var rows [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup

	doorsLockMarks := s.GetAll(uint(page), uint(pagination))

	if len(doorsLockMarks) == 0 {
		return NewAlert(dto, "В процессе заполнения. Попробуйте позже")
	}

	for i := 0; i < len(doorsLockMarks); i++ {
		callbackData := "doorLockModels|doorLockMarkId_" + strconv.FormatUint(doorsLockMarks[i].ID, 10)
		btnText := doorsLockMarks[i].Name
		button := tgbotapi.NewInlineKeyboardButtonData(btnText, callbackData)

		row = append(row, button)
		currentCountInRow := countInRowOptions[rowCountIndex]

		if (len(row) == currentCountInRow) || i == len(doorsLockMarks)-1 {
			rows = append(rows, row)
			row = nil

			rowCountIndex = (rowCountIndex + 1) % len(countInRowOptions)
		}
	}

	paginationDto := PaginationDTO{
		Page:            page,
		QueryCount:      len(doorsLockMarks),
		PaginationCount: pagination,
		CallBack:        "doorLockMarks",
		BackButtonData:  "categories",
		BackButtonText:  "Назад",
	}

	rows = append(rows, getControlPanel(paginationDto))

	keyboard = tgbotapi.NewInlineKeyboardMarkup(rows...)

	var message = messageType.TextWithButtonsMessage{
		Text:    "Выберете производителя замка",
		Bot:     dto.Bot,
		ChatId:  dto.User.TgUserId,
		Buttons: keyboard,
	}

	var screen = DoorLockMarksScreen{OutputMessage: messageType.OutputMessage(message)}

	var baseScreenInterface BaseScreen = screen

	return baseScreenInterface, true
}

func (c DoorLockMarksScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c DoorLockMarksScreen) GetScreenName() string {
	return "doorLockMarks"
}
