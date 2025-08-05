package messageScreens

// func GetScreen(user TgUser.TgUser, bot tgbotapi.BotAPI, name string, db *gorm.DB, filter *uint64, callBack *tgbotapi.CallbackQuery) (baseScreen BaseScreen, clearPrevMessage bool) {
func GetScreen(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	switch dto.ScreenName {

	case "categories":
		return Categories(dto)

		//DOORS
	case "doorLockMarks":
		return DoorLockMarks(dto)
	case "doorLockModels":
		return DoorLockModels(dto)
	case "doorLockModelsDetail":
		return DoorLockModelDetail(dto)
	case "not_found":
		return NotFound(dto)

		//CARS
	case "carMarks":
		return CarMarks(dto)
	case "carModels":
		return CarModels(dto)
	case "carModelDetail":
		return CarModelDetail(dto)

		//COMMANDS
	case "/doorlocks":
		return DoorLockMarks(dto)
	case "/cars":
		return CarMarks(dto)

	}
	return nil, false
}
