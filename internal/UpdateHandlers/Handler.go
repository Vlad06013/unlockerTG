package UpdateHandlers

import (
	"github.com/Vlad06013/unlockerTG.git/internal/messageScreens"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/DoorLockModel"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

var filterMap = make(map[string]string)
var client *TgUser.TgUser
var Bot tgbotapi.BotAPI
var DbConnection *gorm.DB
var UserStorage TgUser.Storage
var screenName string

func SaveLastMessageId(tgMessageId int, clientId uint) {
	if tgMessageId != 0 {
		UserStorage.UpdateLastMessageClient(tgMessageId, clientId)
	}
}

func DeleteLastMessage(chatId int64, messageID int) tgbotapi.Message {
	msg := tgbotapi.NewDeleteMessage(chatId, messageID)
	res, _ := Bot.Send(msg)

	return res
}

func checkNeedAdditionalMessage() {

	if screenName == "doorLockModelsDetail" {
		doorLockModelId, _ := strconv.ParseUint(filterMap["id"], 10, 32)

		s := DoorLockModel.Storage{DB: DbConnection}
		doorLockModel := s.GetById(doorLockModelId)

		filter := map[string]string{
			"doorLockMarkId": strconv.FormatUint(doorLockModel.DoorLockMarkId, 10),
		}
		dto := messageScreens.BaseScreenDTO{
			User:       *client,
			Bot:        Bot,
			DB:         DbConnection,
			ScreenName: "doorLockModels",
			Filter:     filter,
		}

		screenDoorLockModels, _ := messageScreens.GetScreen(dto)

		var sentResult = screenDoorLockModels.GetOutputMessage().Send()
		SaveLastMessageId(sentResult.MessageID, client.ID)
	}
}
func auth() bool {
	if client.TgUserId != 878108763 {
		return false
	}
	return true
}
