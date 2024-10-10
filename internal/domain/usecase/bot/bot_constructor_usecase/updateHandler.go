package bot_constructor_usecase

import (
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bu *BotUseCase) ListenUpdates(bot *entity.BotApi) {
	botApi := bot.Api
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := botApi.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil {
			bu.TextMessageHandler(bot, update.Message)
		}
		if update.CallbackQuery != nil {
			bu.CallbackQueryMessageHandler(bot, update.CallbackQuery)
		}
		//if update.MyChatMember != nil {
		//	ReadMyChatMember(db,)
		//	//	telegram.SetUser(db, update.MyChatMember.From.ID, update.MyChatMember.From.UserName)
		//	//	telegram.SetChatMember(db, *update.MyChatMember, *bot.Bot)
		//}
	}
}
