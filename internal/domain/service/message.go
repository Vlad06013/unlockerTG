package service

import (
	"database/sql"
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
)

type MessageStorage interface {
	GetMessageByID(messageId uint) (*entity.Message, error)
	GetMessageByIDWithFilter(messageId uint, filter string) (*entity.Message, error)
	FirstMessage() (*entity.Message, error)
	QueryForButtons(table string, btnTextKey string, btnCallbackKey string) (*sql.Rows, error)
	GetMessagable(callbackData string) entity.TgMessagable
	QueryForButtonsWithFilter(table string, btnTextKey string, btnCallbackKey string, filterField string, filter string) (*sql.Rows, error)
	GetMessageWithFilter(field string, value any) (*entity.Message, error)
	GetMessagableByNextMessage(toMessageId uint) *entity.TgMessagable
}

type MessageService struct {
	storage MessageStorage
}

func NewMessageService(storage MessageStorage) *MessageService {
	return &MessageService{
		storage: storage,
	}
}

func (s MessageService) GetMessageByID(messageId uint) (*entity.Message, error) {
	message, err := s.storage.GetMessageByID(messageId)
	if err == nil {
		if message.Keyboard.TableName != "" {
			message.Keyboard.QueryButtons, _ = s.QueryForButtons(
				message.Keyboard.TableName,
				message.Keyboard.KeyToButtonText,
				message.Keyboard.KeyToButtonCallbackData,
			)
		}
	}
	return message, err

}

func (s MessageService) GetMessageByIDWithFilter(messageId uint, filter string) (*entity.Message, error) {
	message, err := s.storage.GetMessageByIDWithFilter(messageId, filter)
	return message, err
}

func (s MessageService) QueryForButtonsWithFilter(
	table string,
	btnTextKey string,
	btnCallbackKey string,
	filterField string,
	filter string,
) (*sql.Rows, error) {
	return s.storage.QueryForButtonsWithFilter(table, btnTextKey, btnCallbackKey, filterField, filter)
}

func (s MessageService) FirstMessage() (*entity.Message, error) {
	return s.storage.FirstMessage()
}

func (s MessageService) QueryForButtons(table string, btnTextKey string, btnCallbackKey string) (*sql.Rows, error) {
	return s.storage.QueryForButtons(table, btnTextKey, btnCallbackKey)
}
func (s MessageService) GetMessagable(callbackData string) entity.TgMessagable {
	message := s.storage.GetMessagable(callbackData)
	if message.ToMessage.Keyboard.TableName != "" {
		message.ToMessage.Keyboard.QueryButtons, _ = s.QueryForButtons(
			message.ToMessage.Keyboard.TableName,
			message.ToMessage.Keyboard.KeyToButtonText,
			message.ToMessage.Keyboard.KeyToButtonCallbackData,
		)
	}
	return message
}
func (s MessageService) GetMessageWithFilter(field string, value any) (*entity.Message, error) {
	return s.storage.GetMessageWithFilter(field, value)
}
func (s MessageService) GetMessagableByNextMessage(toMessageId uint) *entity.TgMessagable {
	return s.storage.GetMessagableByNextMessage(toMessageId)
}
