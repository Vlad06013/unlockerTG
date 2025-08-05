package UpdateHandlers

import (
	"github.com/Vlad06013/unlockerTG.git/internal/messageScreens"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func MessageHandler(message *tgbotapi.Message, db *gorm.DB, bot tgbotapi.BotAPI) {
	s := TgUser.Storage{DB: db}
	client := s.InitClient(message.From.ID, message.From.UserName)

	var screen messageScreens.BaseScreen = nil

	if client.LastScreen == nil || *client.LastScreen == "" {
		screen = messageScreens.GetScreen(*client, bot, "categories", db, nil)
	} else {
		screen = messageScreens.GetScreen(*client, bot, message.Text, db, nil)
	}

	if screen == nil {
		screen = messageScreens.GetScreen(*client, bot, "not_found", db, nil)

		//panic("Screen object not found")
	}

	var sentResult = screen.GetOutputMessage().Send()

	SaveLastScreen(s, sentResult.MessageID, client.ID, screen.GetScreenName())
}
