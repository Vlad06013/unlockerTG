package messageScreens

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
	case "/doorlocks":
		return NewDoorLockMarks(dto)
	case "/cars":
		return NewCarMarks(dto)

	}
	return nil, false
}
