package UpdateHandlers

import (
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SaveLastMessageId(s TgUser.Storage, tgMessageId int, clientId uint, screenName string) {
	if tgMessageId != 0 {
		s.UpdateLastMessageClient(tgMessageId, clientId)
	}
}

func DeleteLastMessage(chatId int64, messageID int, bot tgbotapi.BotAPI) tgbotapi.Message {
	msg := tgbotapi.NewDeleteMessage(chatId, messageID)
	res, _ := bot.Send(msg)

	return res
}
