package UpdateHandlers

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/messageScreens"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func CallBackHandler(callBack *tgbotapi.CallbackQuery) {

	client = UserStorage.InitClient(callBack.From.ID, callBack.From.UserName)

	if auth() == false {
		return
	}
	var screen, clearPrevMessage = checkScreen(callBack)
	var sentResult = screen.GetOutputMessage().Send()

	if client.LastTgMessageId != nil && clearPrevMessage == true {
		DeleteLastMessage(callBack.From.ID, *client.LastTgMessageId)
	}
	SaveLastMessageId(sentResult.MessageID, client.ID)
	checkNeedAdditionalMessage()
}

func checkScreen(callBack *tgbotapi.CallbackQuery) (baseScreen messageScreens.BaseScreen, clearPrevMessage bool) {

	setFilterFromCallBack(callBack)

	var screen messageScreens.BaseScreen = nil

	dto := messageScreens.BaseScreenDTO{
		User:       *client,
		Bot:        Bot,
		DB:         DbConnection,
		Filter:     filterMap,
		ScreenName: screenName,
	}

	screen, clearPrevMessage = messageScreens.GetScreen(dto)

	if screen == nil {
		dto.ScreenName = "not_found"
		screen, clearPrevMessage = messageScreens.GetScreen(dto)
		fmt.Println("Не найден экран " + screenName)
	}

	return screen, clearPrevMessage
}

func setFilterFromCallBack(callBack *tgbotapi.CallbackQuery) {
	res := strings.Split(callBack.Data, "|")
	data := res[0]

	if len(res) > 1 {
		for _, v := range res[1:] {
			filterRes := strings.Split(v, "_")
			filterMap[filterRes[0]] = filterRes[1]
		}
	}
	filterMap["callback_id"] = callBack.ID
	screenName = data
}
