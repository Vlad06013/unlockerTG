package tgBotApi

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/UpdateHandlers"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/Setting"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"log"
)

func Listen(db *gorm.DB) {

	s := Setting.Storage{DB: db}
	setting := s.GetBySlug("bot_token")

	botApi, err := tgbotapi.NewBotAPI(setting.Value)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", botApi.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := botApi.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			UpdateHandlers.MessageHandler(update.Message, db, *botApi)
		}
		if update.CallbackQuery != nil {
			UpdateHandlers.CallBackHandler(update.CallbackQuery, db, *botApi)
		}
		if update.MyChatMember != nil {
			fmt.Println(update.MyChatMember)
		}

	}
}
