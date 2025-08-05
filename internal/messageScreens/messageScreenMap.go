package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func GetScreen(user TgUser.TgUser, bot tgbotapi.BotAPI, name string, db *gorm.DB, filter *uint64) BaseScreen {
	switch name {
	case "categories":
		return Categories(user, bot)
	case "doorLockMarks":
		return DoorLockMarks(user, bot, db)
	case "doorLockModels":
		return DoorLockModels(user, bot, db, filter)
	case "doorLockModelsDetail":
		return DoorLockModelDetail(user, bot, db, filter)
	}
	return nil
}
