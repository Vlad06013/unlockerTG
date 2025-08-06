package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

type BaseScreen interface {
	GetOutputMessage() messageType.OutputMessage
	GetScreenName() string
}

type BaseScreenDTO struct {
	User       TgUser.TgUser
	Bot        tgbotapi.BotAPI
	DB         *gorm.DB
	ScreenName string
	Filter     map[string]string
}
