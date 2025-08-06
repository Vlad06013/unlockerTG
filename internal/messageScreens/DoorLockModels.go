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
	markId, _ := strconv.ParseUint(dto.Filter["doorLockMarkId"], 10, 32)

	s := DoorLockModel.Storage{DB: dto.DB}

	doorsLockModels := s.GetByMarkId(markId, uint(page), uint(pagination))

	if len(doorsLockModels) == 0 {
		return NewAlert(dto, "В процессе заполнения. Попробуйте позже")
	}

	bcDTO := ButtonConstructorDTO{
		FieldForText:     "Name",
		PrefixCallback:   "doorLockModelsDetail|id",
		FieldForCallback: "ID",
		Entities:         doorsLockModels,
	}
	rows := GenerateButtons(bcDTO)

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

	var screen = DoorLockModelsScreen{OutputMessage: messageType.OutputMessage(message)}

	var baseScreenInterface BaseScreen = screen

	return baseScreenInterface, true
}

func (c DoorLockModelsScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c DoorLockModelsScreen) GetScreenName() string {
	return "doorLockModels"
}
