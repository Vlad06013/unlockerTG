package bot_constructor_usecase

import (
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

type CallbackParsed struct {
	Id           string
	CallBackData string
	Filter       *string
	Pointer      string
	PointerID    string
}

func (bu *BotUseCase) GenerateAnswerOnTextMessage(user *entity.TgUser) (*entity.Message, *entity.Message) {

	var lastMessage, nextMessage *entity.Message

	if user.BotHistory.LastTGMessageId != 0 {
		lastMessage, _ = bu.messageService.GetMessageByID(user.BotHistory.LastMessageId)
		if lastMessage != nil {
			nextMessage, _ = bu.messageService.GetMessageByID(lastMessage.NextMessageId)
		}
	} else {
		firstMessage, _ := bu.messageService.FirstMessage()
		if firstMessage != nil {
			nextMessage = firstMessage
		}
	}

	return lastMessage, nextMessage
}

func (bu *BotUseCase) GenerateAnswerOnCallbackMessage(user *entity.TgUser, callback *tgbotapi.CallbackQuery) (*entity.Message, CallbackParsed) {

	//var lastMessage, nextMessage *entity.Message
	var nextMessage *entity.Message
	var callbackParsed CallbackParsed

	//if user.BotHistory.LastMessageId != 0 {
	//	lastMessage, _ = bu.messageService.GetMessageByID(user.BotHistory.LastMessageId)

	//if lastMessage != nil {
	//	nextMessage, _ = bu.messageService.GetMessageByID(lastMessage.NextMessageId)
	//}
	//} else {
	//	firstMessage, _ := bu.messageService.FirstMessage()
	//	if firstMessage != nil {
	//		nextMessage = firstMessage
	//	}
	//}

	if strings.Contains(callback.Data, "_") {
		nextMessage, callbackParsed = bu.parseCallback(callback)
	}

	return nextMessage, callbackParsed
	//return lastMessage, nextMessage
}

func (bu *BotUseCase) parseCallback(cb *tgbotapi.CallbackQuery) (*entity.Message, CallbackParsed) {

	data := strings.Split(cb.Data, "/")
	params := strings.Split(data[0], "_")

	cbParsed := CallbackParsed{
		Id:           cb.ID,
		Pointer:      params[0],
		PointerID:    params[1],
		CallBackData: cb.Data,
	}

	if len(data) > 1 {
		paramsFilter := strings.Split(data[1], "_")

		if paramsFilter[0] == "filter" {
			cbParsed.Filter = &paramsFilter[1]
		}
	}

	var nextMessage *entity.Message
	if cbParsed.Pointer == "query" {
		nextMessage = bu.parseQueryBtn(cbParsed)
	}
	if cbParsed.Pointer == "mess" {
		nextMessage = bu.parseMessBtn(&cbParsed)
	}
	if cbParsed.Pointer == "alert" {
		nextMessage = nil
	}

	return nextMessage, cbParsed
}

func (bu *BotUseCase) parseQueryBtn(cbParsed CallbackParsed) *entity.Message {
	messagable := bu.messageService.GetMessagable(cbParsed.CallBackData)
	nextMessage := messagable.ToMessage

	return nextMessage
}

func (bu *BotUseCase) parseMessBtn(cbParsed *CallbackParsed) *entity.Message {
	messageId, _ := strconv.Atoi(cbParsed.PointerID)
	var nextMessage *entity.Message

	if cbParsed.Filter != nil {
		nextMessage, _ = bu.messageService.GetMessageByIDWithFilter(uint(messageId), *cbParsed.Filter)

	} else {
		nextMessage, _ = bu.messageService.GetMessageByID(uint(messageId))
	}
	return nextMessage
}
