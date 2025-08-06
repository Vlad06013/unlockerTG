package messageScreens

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

var pagination = 40
var countInRowOptions = []int{4, 3}
var rowCountIndex = 0

func getControlPanel(dto PaginationDTO) []tgbotapi.InlineKeyboardButton {

	var controlRow []tgbotapi.InlineKeyboardButton

	backBtn := tgbotapi.NewInlineKeyboardButtonData(dto.BackButtonText, dto.BackButtonData)
	nextPageBtn := tgbotapi.NewInlineKeyboardButtonData(">>", dto.CallBack+"|page_"+strconv.FormatUint(dto.Page+1, 10))
	previousPageBtn := tgbotapi.NewInlineKeyboardButtonData("<<", dto.CallBack+"|page_"+strconv.FormatUint(dto.Page-1, 10))

	if dto.Page > 0 {
		controlRow = append(controlRow, previousPageBtn)
	}

	controlRow = append(controlRow, backBtn)

	if dto.QueryCount == dto.PaginationCount {
		controlRow = append(controlRow, nextPageBtn)
	}
	return controlRow
}
