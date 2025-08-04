package tgBotApi

import (
	"fmt"
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

			//fmt.Println(update.Message.Text)
			//var message = messageType.TextMessage{
			//	Text:   update.Message.Text,
			//	Bot:    *botApi,
			//	ChatId: update.Message.Chat.ID,
			//}
			//s := "1"
			//var keyboard tgbotapi.InlineKeyboardMarkup
			//rows := make([][]tgbotapi.InlineKeyboardButton, 2)
			//rows[0] = tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.InlineKeyboardButton{
			//		Text:         "Добавить новый домен",
			//		CallbackData: &s,
			//	},
			//	tgbotapi.InlineKeyboardButton{
			//		Text:         "В кабинет",
			//		CallbackData: &s,
			//	},
			//)
			//rows[1] = tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.InlineKeyboardButton{
			//		Text:         "Добавить новый домен",
			//		CallbackData: &s,
			//	},
			//	tgbotapi.InlineKeyboardButton{
			//		Text:         "В кабинет",
			//		CallbackData: &s,
			//	},
			//)
			//keyboard = tgbotapi.NewInlineKeyboardMarkup(rows...)
			//var message = messageType.TextWithButtonsMessage{
			//	Text:    update.Message.Text,
			//	Bot:     *botApi,
			//	ChatId:  update.Message.Chat.ID,
			//	Buttons: keyboard,
			//}
			//var message = messageType.AlertMessage{
			//	Text:   update.Message.Text,
			//	Bot:    *botApi,
			//	CallBackID: update.Message.Chat.ID,
			//}
			//var output OutputMessage = message
			//output.Send()
			//TextMessage(update.Message, *botApi, conn)
		}
		if update.CallbackQuery != nil {
			UpdateHandlers.CallBackHandler(update.CallbackQuery, db, *botApi)

			fmt.Println(update.CallbackQuery.ID)
			//var message = messageType.AlertMessage{
			//	Text:       "ssss",
			//	Bot:        *botApi,
			//	CallBackID: update.CallbackQuery.ID,
			//}
			//var output messageType.OutputMessage = message
			//output.Send()

			//CallBackQuery(update.CallbackQuery, *botApi, conn)
			//bu.CallbackQueryMessageHandler(bot, update.CallbackQuery)
		}
		//if update.MyChatMember != nil {
		//	ReadMyChatMember(db,)
		//	//	telegram.SetUser(db, update.MyChatMember.From.ID, update.MyChatMember.From.UserName)
		//	//	telegram.SetChatMember(db, *update.MyChatMember, *bot.Bot)
		//}
	}
}
