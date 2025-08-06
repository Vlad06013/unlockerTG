package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarModel"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type CarModelsScreen struct {
	OutputMessage messageType.OutputMessage
}

func NewCarModels(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	id, _ := strconv.ParseUint(dto.Filter["id"], 10, 32)
	page, _ := strconv.ParseUint(dto.Filter["page"], 10, 32)
	pagination := 10
	countInRow := 4

	var row []tgbotapi.InlineKeyboardButton
	var rows [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup

	s := CarModel.Storage{DB: dto.DB}
	carModels := s.GetByMarkId(id, uint(page), uint(pagination))

	if len(carModels) == 0 {
		return NewAlert(dto, "Не найдено.")
	}

	for i := 0; i < len(carModels); i++ {
		callbackData := "carModelDetail|id_" + strconv.FormatUint(uint64(carModels[i].ID), 10)

		btnText := carModels[i].Name
		button := tgbotapi.NewInlineKeyboardButtonData(btnText, callbackData)

		row = append(row, button)

		if (i+1)%countInRow == 0 || i == len(carModels)-1 {
			rows = append(rows, row)
			row = nil
		}
	}

	rows = append(rows, getControlPanel(page, len(carModels), pagination, id))
	keyboard = tgbotapi.NewInlineKeyboardMarkup(rows...)

	var message = messageType.TextWithButtonsMessage{
		Text:    "Выберете модель авто",
		Bot:     dto.Bot,
		ChatId:  dto.User.TgUserId,
		Buttons: keyboard,
	}

	var baseScreenInterface BaseScreen = DoorLockMarksScreen{OutputMessage: messageType.OutputMessage(message)}

	return baseScreenInterface, true
}

func getControlPanel(page uint64, len int, pagination int, idCarMark uint64) []tgbotapi.InlineKeyboardButton {

	var controlRow []tgbotapi.InlineKeyboardButton

	backBtn := tgbotapi.NewInlineKeyboardButtonData("Назад", "carMarks")
	nextPageBtn := tgbotapi.NewInlineKeyboardButtonData(">>", "carModels|id_"+strconv.FormatUint(uint64(idCarMark), 10)+"|page_"+strconv.FormatUint(uint64(page+1), 10))
	previousPageBtn := tgbotapi.NewInlineKeyboardButtonData("<<", "carModels|id_"+strconv.FormatUint(uint64(idCarMark), 10)+"|page_"+strconv.FormatUint(uint64(page-1), 10))

	if page > 0 {
		controlRow = append(controlRow, previousPageBtn)
	}

	controlRow = append(controlRow, backBtn)

	if len == pagination {
		controlRow = append(controlRow, nextPageBtn)
	}
	return controlRow
}

func (c CarModelsScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c CarModelsScreen) GetScreenName() string {
	return "carModels"
}
