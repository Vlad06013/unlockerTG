package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

type BaseScreen interface {
	GetOutputMessage() messageType.OutputMessage
	GetScreenName() string
}

type BaseScreenDTO struct {
	User       TgUser.TgUser
	Bot        tgbotapi.BotAPI
	DB         *gorm.DB
	ScreenName string
	Filter     map[string]string
}
type PaginationDTO struct {
	Page            uint64
	QueryCount      int
	PaginationCount int
	CallBack        string
	BackButtonData  string
	BackButtonText  string
}

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
