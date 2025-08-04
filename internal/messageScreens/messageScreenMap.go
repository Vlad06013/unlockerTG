package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func Map(user TgUser.TgUser, bot tgbotapi.BotAPI, db *gorm.DB, filter *uint64) map[string]BaseScreen {

	var screenMap = map[string]BaseScreen{
		"categories":     Categories(user, bot),
		"doorLockMarks":  DoorLockMarks(user, bot, db),
		"doorLockModels": DoorLockModels(user, bot, db, filter),
	}
	return screenMap
}

func GetScreen(user TgUser.TgUser, bot tgbotapi.BotAPI, name string, db *gorm.DB, filter *uint64) BaseScreen {
	var screenMap = Map(user, bot, db, filter)
	return screenMap[name]
}
