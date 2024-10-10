package bot_constructor_usecase

import (
	"github.com/Vlad06013/unlockerTG.git/internal/controller/api/telegram"
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
	"github.com/Vlad06013/unlockerTG.git/internal/domain/usecase/bot/bot_constructor_usecase/messageTypes/defaultMessage"
	"github.com/Vlad06013/unlockerTG.git/internal/domain/usecase/bot/bot_constructor_usecase/messageTypes/queryBtnMessage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (bu BotUseCase) SendAnswer() *tgbotapi.Message {
	result := bu.output.SendMessage()
	sendedMessage := bu.messageConstruct.GetMessage()
	if result != nil && bu.messageConstruct != nil && sendedMessage.ID != 0 {
		bu.SaveLastMessage(sendedMessage.ID, uint(result.MessageID))
	}
	if bu.CheckSendNextMessage() {
		bu.sendNextAnswer()
	}

	return result

}

func (bu BotUseCase) sendNextAnswer() {
	user := bu.userService.User
	_, nextMessage := bu.GenerateAnswerOnTextMessage(user)
	if nextMessage != nil {
		messageConstruct := bu.generateConstruct(user, nextMessage)
		if messageConstruct != nil {
			bu.messageConstruct = messageConstruct
			bu.output = &telegram.Output{
				ChatId:      bu.messageConstruct.Chat(),
				Text:        bu.messageConstruct.TextMessage(),
				TypeMessage: bu.messageConstruct.TypeMessage(),
				Buttons:     bu.messageConstruct.ButtonsMessage(),
				Bot:         bu.output.Bot,
			}
			bu.SendAnswer()
		}
	}

}
func (bu BotUseCase) generateConstruct(user *entity.TgUser, message *entity.Message) entity.Constructable {

	var messageConstruct entity.Constructable
	if message.Keyboard.TableName != "" {
		construct := queryBtnMessage.BtnQueryMessageConstructor{
			ChatId:         user.TgUserId,
			Text:           message.Text,
			Type:           message.Type,
			Keyboard:       message.Keyboard,
			Message:        message,
			MessageService: &bu.messageService,
		}
		messageConstruct = queryBtnMessage.NewMessage(construct)
	} else {
		construct := defaultMessage.DefaultMessageConstructor{
			ChatId:   user.TgUserId,
			Text:     message.Text,
			Type:     message.Type,
			Keyboard: message.Keyboard,
			Message:  message,
		}
		messageConstruct = defaultMessage.NewMessage(construct)
	}

	return messageConstruct
}
