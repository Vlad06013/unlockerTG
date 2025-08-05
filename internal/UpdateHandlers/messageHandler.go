package UpdateHandlers

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/messageScreens"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

func MessageHandler(message *tgbotapi.Message, db *gorm.DB, bot tgbotapi.BotAPI) {
	s := TgUser.Storage{DB: db}
	client := s.InitClient(message.From.ID, message.From.UserName)

	var screen messageScreens.BaseScreen = nil
	dto := messageScreens.BaseScreenDTO{
		User:   *client,
		Bot:    bot,
		DB:     db,
		Filter: nil,
	}
	if client.LastScreen == nil || *client.LastScreen == "" {
		dto.ScreenName = "categories"
		screen, _ = messageScreens.GetScreen(dto)
	} else {
		dto.ScreenName = message.Text
		screen, _ = messageScreens.GetScreen(dto)
	}

	if screen == nil {
		dto.ScreenName = "not_found"
		screen, _ = messageScreens.GetScreen(dto)
		fmt.Println("Не найден экран " + message.Text)
	}

	var sentResult = screen.GetOutputMessage().Send()

	if client.LastTgMessageId != nil {
		DeleteLastMessage(message.From.ID, *client.LastTgMessageId, bot)
	}
	SaveLastScreen(s, sentResult.MessageID, client.ID, screen.GetScreenName())
}
