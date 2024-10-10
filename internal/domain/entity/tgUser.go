package entity

type TgUser struct {
	ID         uint   `json:"id" gorm:"primary_key;column:id"`
	UserId     uint   `json:"user_id" gorm:"column:user_id;default:null"`
	Name       string `json:"name" gorm:"column:name;default:noname"`
	Email      string `json:"email" gorm:"column:email;default:null"`
	Phone      string `json:"phone" gorm:"column:phone;default:null"`
	TgUserId   int64  `json:"tg_user_id" gorm:"column:tg_user_id;unique"`
	TgUserName string `json:"tg_user_name" gorm:"column:tg_user_name"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at"`

	BotHistory *TgUserMessageHistory
}

type TgUserMessageHistory struct {
	ID              uint    `json:"id" gorm:"primary_key;column:id"`
	TgUserId        uint    `json:"tg_user_id" gorm:"column:tg_user_id"`
	BotId           uint    `json:"bot_id" gorm:"column:bot_id"`
	LastMessageId   uint    `json:"last_message_id" gorm:"foreignKey:id;default:null;OnDelete:SET NULL;"`
	LastTGMessageId uint    `json:"last_tg_message_id" gorm:"column:last_tg_message_id;default:null;OnDelete:SET NULL;"`
	LastQueryFilter string  `json:"last_query_filter" gorm:"column:last_query_filter;default:null;OnDelete:SET NULL;"`
	LastMessage     Message `json:"last_message" gorm:"foreignKey:LastMessageId;default:null"`
}
