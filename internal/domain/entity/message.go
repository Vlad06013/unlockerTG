package entity

const modelNameSpace = "Valibool\\TelegramConstruct\\Models\\Message"

type Message struct {
	ID               uint   `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	Name             string `json:"name" gorm:"column:name"`
	Type             string `json:"type" gorm:"column:type"`
	Text             string `json:"text" gorm:"column:text"`
	FirstMessage     bool   `json:"first_message" gorm:"column:first_message"`
	WaitInput        string `json:"wait_input" gorm:"column:wait_input"`
	NeedConfirmation bool   `json:"need_confirmation" gorm:"column:need_confirmation"`
	BotId            uint   `json:"bot_id" gorm:"column:bot_id;foreign_key:bot_id"`
	NextMessageId    uint   `json:"next_message_id" gorm:"column:next_message_id;foreign_key:next_message_id"`

	Bot         *Bot     `json:"bot" gorm:"foreignKey:BotId;default:null"`
	NextMessage *Message `json:"next_message" gorm:"foreignKey:NextMessageId;default:null"`
	Keyboard    Keyboard
}

type TgMessagable struct {
	ID            uint     `gorm:"AUTO_INCREMENT;primary_key;column:id"`
	FromMessageId uint     `gorm:"column:from_message_id;foreign_key:from_message_id"`
	ToMessageId   uint     `gorm:"column:to_message_id;foreign_key:to_message_id"`
	CallbackData  string   `gorm:"column:callback_data"`
	FromMessage   *Message `gorm:"foreignKey:FromMessageId;default:null"`
	ToMessage     *Message `gorm:"foreignKey:ToMessageId;default:null"`
}
