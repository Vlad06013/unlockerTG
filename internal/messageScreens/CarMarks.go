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

	s := CarMark.Storage{DB: dto.DB}

	carMarks := s.GetAll(uint(page), uint(pagination))

	bcDTO := ButtonConstructorDTO{
		FieldForText:     "Name",
		PrefixCallback:   "carModels|id",
		FieldForCallback: "ID",
		Entities:         carMarks,
	}
	rows := GenerateButtons(bcDTO)

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

	var screen = CarMarksScreen{OutputMessage: messageType.OutputMessage(message)}
	var baseScreenInterface BaseScreen = screen

	return baseScreenInterface, true
}

func (c CarMarksScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c CarMarksScreen) GetScreenName() string {
	return "carMarks"
}
