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

func CarModels(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	id, _ := strconv.ParseUint(dto.Filter["id"], 10, 32)
	page, _ := strconv.ParseUint(dto.Filter["page"], 10, 32)
	pagination := 10

	backBtn := tgbotapi.NewInlineKeyboardButtonData("Назад", "carMarks")
	nextPageBtn := tgbotapi.NewInlineKeyboardButtonData(">>", "carModels|id_"+strconv.FormatUint(uint64(id), 10)+"|page_"+strconv.FormatUint(uint64(page+1), 10))
	previousPageBtn := tgbotapi.NewInlineKeyboardButtonData("<<", "carModels|id_"+strconv.FormatUint(uint64(id), 10)+"|page_"+strconv.FormatUint(uint64(page-1), 10))

	s := CarModel.Storage{DB: dto.DB}
	var buttons [][]tgbotapi.InlineKeyboardButton
	var keyboard tgbotapi.InlineKeyboardMarkup

	carModels := s.GetByMarkId(id, uint(page), uint(pagination))
	text := "Выберете модель авто"

	if len(carModels) == 0 {
		text = "Модели отсутствуют"
	}
	var row []tgbotapi.InlineKeyboardButton
	var controlRow []tgbotapi.InlineKeyboardButton
	var rows [][]tgbotapi.InlineKeyboardButton
	countInRow := 4

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

	if page > 0 {
		controlRow = append(controlRow, previousPageBtn)
	}

	controlRow = append(controlRow, backBtn)

	if len(carModels) == pagination {
		controlRow = append(controlRow, nextPageBtn)
	}

	rows = append(rows, controlRow)
	keyboard = tgbotapi.NewInlineKeyboardMarkup(rows...)

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

func (c CarModelsScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c CarModelsScreen) GetScreenName() string {
	return "carModels"
}
