package UpdateHandlers

import (
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
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
