package bot_constructor_usecase

import (
	"github.com/Vlad06013/unlockerTG.git/internal/controller/api/telegram"
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bu BotUseCase) CallbackQueryMessageHandler(bot *entity.BotApi, callback *tgbotapi.CallbackQuery) {

	user := bu.userService.InitUser(callback.From.ID, callback.From.UserName, bot.Bot)

	//maybe layer validated and access
	nextMessage, _ := bu.GenerateAnswerOnCallbackMessage(user, callback)

	if nextMessage != nil {
		messageConstruct := bu.generateConstruct(user, nextMessage)
		if messageConstruct != nil {
			bu.messageConstruct = messageConstruct
			bu.output = &telegram.Output{
				ChatId:      bu.messageConstruct.Chat(),
				Text:        bu.messageConstruct.TextMessage(),
				TypeMessage: bu.messageConstruct.TypeMessage(),
				Buttons:     bu.messageConstruct.ButtonsMessage(),
				Bot:         bot.Api,
			}

			go bu.output.DeleteMessage(int(bu.userService.User.BotHistory.LastTGMessageId))
			bu.SendAnswer()

		}
	}
}
