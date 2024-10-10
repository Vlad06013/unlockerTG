package postgres

import (
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
	"github.com/jinzhu/gorm"
	"time"
)

type userStorage struct {
	*gorm.DB
}

func NewUserStorage(pg *gorm.DB) *userStorage {
	return &userStorage{pg}
}

func (r *userStorage) GetUserByTGID(tgID int64) (*entity.TgUser, error) {
	var user entity.TgUser
	var history entity.TgUserMessageHistory
	if err := r.First(&user, "tg_user_id = ?", tgID).Error; err != nil {
		return nil, err
	}
	if err := r.First(&history, "tg_user_id = ?", user.ID).Error; err == nil {
		user.BotHistory = &history
	}

	return &user, nil
}

func (r *userStorage) CreateUser(u entity.TgUser) *entity.TgUser {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	user := entity.TgUser{
		TgUserId:   u.TgUserId,
		TgUserName: u.Name,
		Name:       u.Name,
		CreatedAt:  dateTime,
		UpdatedAt:  dateTime,
	}
	r.Create(&user)
	return &user
}

func (r *userStorage) CreateUserHistory(TgUserId uint, botId uint) *entity.TgUserMessageHistory {

	history := entity.TgUserMessageHistory{
		BotId:    botId,
		TgUserId: TgUserId,
	}
	r.Create(&history)
	return &history
}

func (r *userStorage) GetHistoryByBotId(tgUserId int64, botId uint) (entity.TgUserMessageHistory, error) {
	var history entity.TgUserMessageHistory

	err := r.Table("tg_user_message_histories").
		Preload("LastMessage").
		Where("bot_id = ?", &botId).Where("tg_user_id = ?", &tgUserId).First(&history).Error

	return history, err
}

func (r *userStorage) UpdateHistory(Id uint, updates *entity.TgUserMessageHistory) *entity.TgUserMessageHistory {

	history := entity.TgUserMessageHistory{}

	r.First(&history, Id)
	history.LastMessageId = updates.LastMessageId
	history.LastTGMessageId = updates.LastTGMessageId
	r.Save(&history)
	r.Preload("LastMessage").Preload("LastMessage.Keyboard").Preload("LastMessage.Keyboard.Buttons").First(&history)
	return &history
}
