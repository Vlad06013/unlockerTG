package messageScreens

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/SearchModule"
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
	findByText(dto.ScreenName, dto)
	return nil, false
}

func findByText(screenName string, dto BaseScreenDTO) {
	SearchModule.SearchString = screenName
	SearchModule.DbConnection = dto.DB
	findResult := SearchModule.Find()
	fmt.Println(findResult)
}
