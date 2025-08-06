package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
)

type NotFoundAlert struct {
	OutputMessage messageType.OutputMessage
}

func NewAlert(dto BaseScreenDTO, text string) (baseScreen BaseScreen, clearPrevMessage bool) {

	var message = messageType.AlertMessage{
		Text:       text,
		Bot:        dto.Bot,
		CallBackID: dto.Filter["callback_id"],
	}
	var baseScreenInterface BaseScreen = NotFoundAlert{OutputMessage: message}
	return baseScreenInterface, false

}

func (c NotFoundAlert) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c NotFoundAlert) GetScreenName() string {
	return "alert"
}
