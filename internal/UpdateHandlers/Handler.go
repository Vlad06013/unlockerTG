package UpdateHandlers

import (
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SaveLastScreen(s TgUser.Storage, tgMessageId int, clientId uint, screenName string) {
	s.UpdateLastMessageClient(tgMessageId, clientId)
	s.SaveLastScreenWithFilter(screenName, nil, clientId)
}

func DeleteLastMessage(chatId int64, messageID int, bot tgbotapi.BotAPI) tgbotapi.Message {
	msg := tgbotapi.NewDeleteMessage(chatId, messageID)
	res, _ := bot.Send(msg)

	return res
}
