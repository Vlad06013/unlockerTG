package UpdateHandlers

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/messageScreens"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strings"
)

func CallBackHandler(callBack *tgbotapi.CallbackQuery, db *gorm.DB, bot tgbotapi.BotAPI) {
	s := TgUser.Storage{DB: db}
	client := s.InitClient(callBack.From.ID, callBack.From.UserName)

	var screen, clearPrevMessage = checkScreen(callBack, client, bot, db)

	if screen != nil {
		var sentResult = screen.GetOutputMessage().Send()

		if client.LastTgMessageId != nil && clearPrevMessage == true {
			DeleteLastMessage(callBack.From.ID, *client.LastTgMessageId, bot)
		}
		SaveLastMessageId(s, sentResult.MessageID, client.ID, screen.GetScreenName())
	}
}

func checkScreen(callBack *tgbotapi.CallbackQuery, client *TgUser.TgUser, bot tgbotapi.BotAPI, db *gorm.DB) (baseScreen messageScreens.BaseScreen, clearPrevMessage bool) {

	dataParsed, filter := parseCallBack(callBack)
	data := dataParsed
	var screen messageScreens.BaseScreen = nil

	dto := messageScreens.BaseScreenDTO{
		User:   *client,
		Bot:    bot,
		DB:     db,
		Filter: filter,
	}
	dto.ScreenName = data
	screen, clearPrevMessage = messageScreens.GetScreen(dto)

	if screen == nil {
		dto.ScreenName = "not_found"
		screen, clearPrevMessage = messageScreens.GetScreen(dto)
		fmt.Println("Не найден экран " + data)
	}

	return screen, clearPrevMessage
}

func parseCallBack(callBack *tgbotapi.CallbackQuery) (string, map[string]string) {
	res := strings.Split(callBack.Data, "|")

	var filterMap = make(map[string]string)
	data := res[0]

	if len(res) > 1 {
		for _, v := range res[1:] {
			filterRes := strings.Split(v, "_")
			filterMap[filterRes[0]] = filterRes[1]
		}
	}
	filterMap["callback_id"] = callBack.ID

	return data, filterMap
}
