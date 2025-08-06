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

func NewDoorLockModels(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	page, _ := strconv.ParseUint(dto.Filter["page"], 10, 32)
	pagination := 40
	countInRow := 4

	s := DoorLockModel.Storage{DB: dto.DB}

	var row []tgbotapi.InlineKeyboardButton
	var rows [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup

	markId, _ := strconv.ParseUint(dto.Filter["id"], 10, 32)
	doorsLockModels := s.GetByMarkId(markId, uint(page), uint(pagination))

	if len(doorsLockModels) == 0 {
		return NewAlert(dto, "В процессе заполнения. Попробуйте позже")
	}

	for i := 0; i < len(doorsLockModels); i++ {
		callbackData := "doorLockModelsDetail|id_" + strconv.FormatUint(uint64(doorsLockModels[i].ID), 10)
		btnText := doorsLockModels[i].Name
		button := tgbotapi.NewInlineKeyboardButtonData(btnText, callbackData)

		row = append(row, button)

		if (i+1)%countInRow == 0 || i == len(doorsLockModels)-1 {
			rows = append(rows, row)
			row = nil
		}
	}

	paginationDto := PaginationDTO{
		Page:            page,
		QueryCount:      len(doorsLockModels),
		PaginationCount: pagination,
		CallBack:        "doorLockModels|id_" + strconv.FormatUint(markId, 10),
		BackButtonData:  "doorLockMarks",
		BackButtonText:  "Назад",
	}

	rows = append(rows, getControlPanel(paginationDto))
	keyboard = tgbotapi.NewInlineKeyboardMarkup(rows...)

	var message = messageType.TextWithButtonsMessage{
		Text:    "Выберете модель замка",
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
