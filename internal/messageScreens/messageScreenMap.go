package messageScreens

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/internal/SearchModule"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetScreen(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {

	switch dto.ScreenName {

	case "categories":
		return NewCategoriesScreen(dto)

		//DOORS
	case "doorLockMarks":
		return NewDoorLockMarks(dto)
	case "doorLockModels":
		return NewDoorLockModels(dto)
	case "doorLockModelsDetail":
		return NewDoorLockModelDetail(dto)
	case "not_found":
		return NewNotFound(dto)

		//CARS
	case "carMarks":
		return NewCarMarks(dto)
	case "carModels":
		return NewCarModels(dto)
	case "carModelDetail":
		return NewCarModelDetail(dto)

		//COMMANDS
	case "/start":
		return NewCategoriesScreen(dto)
	case "/doorlocks":
		return NewDoorLockMarks(dto)
	case "/cars":
		return NewCarMarks(dto)
	}

	found, cPrevMess := findByText(dto.ScreenName, dto)

	if found != nil {
		return found, cPrevMess
	}

	return nil, false
}

func findByText(searchString string, dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	SearchModule.SearchString = searchString
	SearchModule.DbConnection = dto.DB
	findResult := SearchModule.Find()

	var rows [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(findResult); i++ {

		switch findResult[i].Type {
		case "carMark":
			bcDTO := ButtonConstructorDTO{
				FieldForText:     "Name",
				PrefixCallback:   "carModels|id",
				FieldForCallback: "ID",
				Entities:         findResult[i].Entities,
			}
			rows = GenerateButtons(bcDTO)
			break

		case "carModel":
			bcDTO := ButtonConstructorDTO{
				FieldForText:     "Name",
				PrefixCallback:   "carModelDetail|id",
				FieldForCallback: "ID",
				Entities:         findResult[i].Entities,
			}
			rows = GenerateButtons(bcDTO)
			break
		}
	}
	keyboard = tgbotapi.NewInlineKeyboardMarkup(rows...)

	var message = messageType.TextWithButtonsMessage{
		Text:    fmt.Sprintf("Найдено по запросу \"%s\"", searchString),
		Bot:     dto.Bot,
		ChatId:  dto.User.TgUserId,
		Buttons: keyboard,
	}

	var screen = CarMarksScreen{OutputMessage: messageType.OutputMessage(message)}
	var baseScreenInterface BaseScreen = screen

	return baseScreenInterface, true
}
