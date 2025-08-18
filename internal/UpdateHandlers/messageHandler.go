package UpdateHandlers

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/messageScreens"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageHandler(message *tgbotapi.Message) {
	client := UserStorage.InitClient(message.From.ID, message.From.UserName)

	if auth() == false {
		return
	}

	var screen messageScreens.BaseScreen = nil
	dto := messageScreens.BaseScreenDTO{
		User:   *client,
		Bot:    Bot,
		DB:     DbConnection,
		Filter: nil,
	}

	dto.ScreenName = message.Text
	screen, _ = messageScreens.GetScreen(dto)

	if screen == nil {
		dto.ScreenName = "not_found"
		screen, _ = messageScreens.GetScreen(dto)
		fmt.Println("Не найден экран " + message.Text)
	}

	var sentResult = screen.GetOutputMessage().Send()

	if client.LastTgMessageId != nil {
		DeleteLastMessage(message.From.ID, *client.LastTgMessageId)
	}
	SaveLastMessageId(sentResult.MessageID, client.ID)
}
