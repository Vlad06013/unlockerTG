package postgres

import (
	"database/sql"
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
	"github.com/jinzhu/gorm"
)

type messageStorage struct {
	*gorm.DB
}

func NewMessageStorage(pg *gorm.DB) *messageStorage {
	return &messageStorage{pg}
}

func (r *messageStorage) GetMessageByIDWithFilter(messageId uint, filter string) (*entity.Message, error) {
	var message entity.Message
	err := r.Preload("Keyboard").Preload("Keyboard.Buttons").First(&message, messageId).Error
	if message.Keyboard.TableName != "" {
		if message.Keyboard.InputFilterField != "" {
			buttons, err := r.QueryForButtonsWithFilter(
				message.Keyboard.TableName,
				message.Keyboard.KeyToButtonText,
				message.Keyboard.KeyToButtonCallbackData,
				message.Keyboard.InputFilterField,
				filter,
			)
			if err == nil {
				message.Keyboard.QueryButtons = buttons
			}
		}

	}

	return &message, err
}
func (r *messageStorage) GetMessageByID(messageId uint) (*entity.Message, error) {
	var message entity.Message
	err := r.Preload("Keyboard").Preload("Keyboard.Buttons").First(&message, messageId).Error

	return &message, err
}

func (r *messageStorage) FirstMessage() (*entity.Message, error) {
	var firstMessage entity.Message
	err := r.Where("first_message = ?", true).Preload("Keyboard").Preload("Keyboard.Buttons").Find(&firstMessage).Error
	return &firstMessage, err
}

func (r *messageStorage) QueryForButtons(table string, btnTextKey string, btnCallbackKey string) (*sql.Rows, error) {
	rows, err := r.Table(table).
		Select([]string{btnTextKey, btnCallbackKey}).Rows()
	return rows, err
}
func (r *messageStorage) QueryForButtonsWithFilter(
	table string,
	btnTextKey string,
	btnCallbackKey string,
	filterField string,
	filter string,
) (*sql.Rows, error) {
	rows, err := r.DB.Table(table).
		Select([]string{btnTextKey, btnCallbackKey}).Where(filterField+"=?", filter).Rows()
	return rows, err
}

func (r *messageStorage) GetMessagable(callbackData string) entity.TgMessagable {

	var messagable entity.TgMessagable
	r.Where("callback_data = ?", callbackData).
		Preload("ToMessage").
		Preload("ToMessage.Keyboard").
		Preload("ToMessage.Keyboard.Buttons").
		First(&messagable)

	return messagable
}

func (r *messageStorage) GetMessageWithFilter(field string, value any) (*entity.Message, error) {
	var message entity.Message
	err := r.Where(field+" = ?", value).Preload("Keyboard").Preload("Keyboard.Buttons").First(&message).Error
	return &message, err
}

func (r *messageStorage) GetMessagableByNextMessage(toMessageId uint) *entity.TgMessagable {
	messagable := entity.TgMessagable{}

	if err := r.Where("to_message_id = ?", toMessageId).
		Preload("FromMessage").
		Preload("FromMessage.Keyboard").
		Preload("FromMessage.Keyboard.Buttons").
		First(&messagable).Error; err != nil {
		return nil
	}

	return &messagable
}
