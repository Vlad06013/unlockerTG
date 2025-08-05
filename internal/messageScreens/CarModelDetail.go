package messageScreens

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarModel"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
)

type CarModelDetailScreen struct {
	OutputMessage messageType.OutputMessage
}

func CarModelDetail(user TgUser.TgUser, bot tgbotapi.BotAPI, db *gorm.DB, filter *uint64, callBack *tgbotapi.CallbackQuery) BaseScreen {
	s := CarModel.Storage{DB: db}
	carModel := s.GetById(*filter)
	fmt.Println(callBack.ID, callBack.Data)

	iconFalse := "❌"
	iconTrue := "✅"
	openInside, openByClockArrowUp := iconFalse, iconFalse

	if carModel.OpenByClockArrowUp {
		openByClockArrowUp = iconTrue
	}

	if carModel.OpenInside {
		openInside = iconTrue
	}

	//max text lenght 200 chars
	var text = carModel.CarMark.Name + " " + carModel.Name + `
	Открыть через салон: ` + openInside + `.
	Открывается по часовой стрелке: ` + openByClockArrowUp + `.
	Транспондер: ` + carModel.Transponder.Name + `.`
	//Профиль: ` + carModel.Profile + `.
	//Подготовка: ` + carModel.Prepare.Name + `.
	//Программирование: ` + carModel.Programming.Name + `.
	//Код: ` + carModel.Code + `.
	//Пульт: ` + carModel.Pult.Name + `.`
	//Описание: ` + carModel.Description + `.`

	var message = messageType.AlertMessage{
		Text:       text,
		Bot:        bot,
		CallBackID: callBack.ID,
	}

	var screen = DoorLockMarksScreen{
		OutputMessage: messageType.OutputMessage(message),
	}

	var baseScreen BaseScreen = screen

	return baseScreen
}

func (c CarModelDetailScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c CarModelDetailScreen) GetScreenName() string {
	return "carModelDetail"
}
