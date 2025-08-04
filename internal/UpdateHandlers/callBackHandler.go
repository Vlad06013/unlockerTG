package UpdateHandlers

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/messageScreens"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"

	//"strconv"
	"strings"
)

func CallBackHandler(callBack *tgbotapi.CallbackQuery, db *gorm.DB, bot tgbotapi.BotAPI) {
	s := TgUser.Storage{DB: db}
	client := s.InitClient(callBack.From.ID, callBack.From.UserName)

	//fmt.Print(callBack.Data)
	var screen = checkScreen(callBack.Data, client, bot, db)

	if screen != nil {
		var sentResult = screen.GetOutputMessage().Send()

		if client.LastTgMessageId != nil {
			DeleteLastMessage(callBack.From.ID, *client.LastTgMessageId, bot)
		}

		SaveLastScreen(s, sentResult.MessageID, client.ID, screen.GetScreenName())
	}
}

func checkScreen(data string, client *TgUser.TgUser, bot tgbotapi.BotAPI, db *gorm.DB) messageScreens.BaseScreen {
	dataParsed, filter := parseCallBack(data)
	data = dataParsed
	var screen messageScreens.BaseScreen = nil

	if client.LastScreen == nil {
		return messageScreens.GetScreen(*client, bot, "categories", db, nil)
	}

	screen = messageScreens.GetScreen(*client, bot, data, db, filter)

	if screen == nil {
		fmt.Println("Не найден экран " + data)
	}

	return screen
}

func parseCallBack(data string) (string, *uint64) {
	res := strings.Split(data, "|")

	var filter uint64
	data = res[0]

	if len(res) > 1 {
		filter, _ = strconv.ParseUint(res[1], 10, 32)
	}
	fmt.Println(filter)
	fmt.Println(data)

	return data, &filter
}
