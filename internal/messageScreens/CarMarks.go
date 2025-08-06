package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarMark"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type CarMarksScreen struct {
	OutputMessage messageType.OutputMessage
}

func NewCarMarks(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {

	page, _ := strconv.ParseUint(dto.Filter["page"], 10, 32)
	pagination := 42
	countInRowOptions := []int{4, 3}
	rowCountIndex := 0

	var row []tgbotapi.InlineKeyboardButton
	var rows [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup
	s := CarMark.Storage{DB: dto.DB}

	carMarks := s.GetAll(uint(page), uint(pagination))

	if len(carMarks) == 0 {
		return NewAlert(dto, "В процессе заполнения. Попробуйте позже")
	}

	for i := 0; i < len(carMarks); i++ {
		callbackData := "carModels|id_" + strconv.FormatUint(carMarks[i].ID, 10)
		btnText := carMarks[i].Name
		button := tgbotapi.NewInlineKeyboardButtonData(btnText, callbackData)

		row = append(row, button)

		currentCountInRow := countInRowOptions[rowCountIndex]

		if (len(row) == currentCountInRow) || i == len(carMarks)-1 {
			rows = append(rows, row)
			row = nil

			rowCountIndex = (rowCountIndex + 1) % len(countInRowOptions)
		}
	}

	paginationDto := PaginationDTO{
		Page:            page,
		QueryCount:      len(carMarks),
		PaginationCount: pagination,
		CallBack:        "carMarks",
		BackButtonData:  "categories",
		BackButtonText:  "Назад",
	}

	rows = append(rows, getControlPanel(paginationDto))
	keyboard = tgbotapi.NewInlineKeyboardMarkup(rows...)

	var message = messageType.TextWithButtonsMessage{
		Text:    "Выберете марку",
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

func (c CarMarksScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c CarMarksScreen) GetScreenName() string {
	return "carMarks"
}
