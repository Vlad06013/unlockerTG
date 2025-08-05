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

		//DOORS
	case "doorLockMarks":
		return DoorLockMarks(user, bot, db)
	case "doorLockModels":
		return DoorLockModels(user, bot, db, filter)
	case "doorLockModelsDetail":
		return DoorLockModelDetail(user, bot, db, filter)
	case "not_found":
		return NotFound(user, bot)

		//CARS
	case "carMarks":
		return CarMarks(user, bot, db)
	case "carModels":
		return CarModels(user, bot, db, filter)

		//COMMANDS
	case "/doorlocks":
		return DoorLockMarks(user, bot, db)
	case "/cars":
		return CarMarks(user, bot, db)
	}
	return nil
}
