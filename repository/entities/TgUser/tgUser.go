package TgUser

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TgUser struct {
	ID              uint   `json:"id" gorm:"primary_key;column:id"`
	TgUserId        int64  `json:"tg_user_id" gorm:"column:tg_user_id;unique"`
	TgUserName      string `json:"tg_user_name" gorm:"column:tg_user_name"`
	LastTgMessageId *int   `json:"last_tg_message_id" gorm:"column:last_tg_message_id"`
	CreatedAt       string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       string `json:"updated_at" gorm:"column:updated_at"`
}

type Storage struct {
	*gorm.DB
}

func (r *Storage) InitClient(tgID int64, name string) *TgUser {

	client, err := r.GetClientByTGID(tgID)
	if err != nil {
		client = r.CreateClient(tgID, name)
	}

	return client
}

func (r *Storage) GetClientByTGID(tgID int64) (*TgUser, error) {
	var client TgUser
	if err := r.First(&client, "tg_user_id = ?", tgID).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *Storage) CreateClient(tgID int64, name string) *TgUser {

	location, _ := time.LoadLocation("Europe/Moscow")
	dateTime := time.Now().In(location).Format("2006-01-02 15:04:05")

	client := TgUser{
		TgUserId:        tgID,
		TgUserName:      name,
		LastTgMessageId: nil,
		CreatedAt:       dateTime,
		UpdatedAt:       dateTime,
	}
	r.Create(&client)
	return &client
}

func (r *Storage) UpdateLastMessageClient(lastTgMessageId int, clientId uint) {
	r.Model(&TgUser{}).Where("id =?", clientId).Update("last_tg_message_id", lastTgMessageId)
}
