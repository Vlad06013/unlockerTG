package tgBotApi

import (
	"github.com/Vlad06013/unlockerTG.git/internal/UpdateHandlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"log"
)

func Listen(db *gorm.DB) {

	botApi, err := tgbotapi.NewBotAPI("8362746471:AAFPTxWcNirckvvCtsNknFaIFSKDWLuZrsM")
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
			//fmt.Println(update.MyChatMember)

			//	ReadMyChatMember(db,)
			//	//	telegram.SetUser(db, update.MyChatMember.From.ID, update.MyChatMember.From.UserName)
			//	//	telegram.SetChatMember(db, *update.MyChatMember, *bot.Bot)
		}
	}
}
