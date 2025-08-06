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

	doorsLockMarks := s.GetAll(uint(page), uint(pagination))

	if len(doorsLockMarks) == 0 {
		return NewAlert(dto, "В процессе заполнения. Попробуйте позже")
	}

	bcDTO := ButtonConstructorDTO{
		FieldForText:     "Name",
		PrefixCallback:   "doorLockModels|doorLockMarkId",
		FieldForCallback: "ID",
		Entities:         doorsLockMarks,
	}
	rows := GenerateButtons(bcDTO)

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
