package messageScreens

import "github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"

type BaseScreen interface {
	GetOutputMessage() messageType.OutputMessage
	GetScreenName() string
}
