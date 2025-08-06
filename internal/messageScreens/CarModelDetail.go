package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarModel"
	"strconv"
)

type CarModelDetailScreen struct {
	OutputMessage messageType.OutputMessage
}

func NewCarModelDetail(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	id, _ := strconv.ParseUint(dto.Filter["id"], 10, 32)

	s := CarModel.Storage{DB: dto.DB}
	carModel := s.GetById(id)

	if carModel == nil {
		return NewAlert(dto, "Не найдено.")
	}
	return NewAlert(dto, getTextCarModel(*carModel))
}

func getTextCarModel(carModel CarModel.CarModel) string {
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
	return carModel.CarMark.Name + " " + carModel.Name + `
	Открыть через салон: ` + openInside + `.
	Открывается по часовой стрелке: ` + openByClockArrowUp + `.
	Транспондер: ` + carModel.Transponder.Name + `.`
	//Профиль: ` + carModel.Profile + `.
	//Подготовка: ` + carModel.Prepare.Name + `.
	//Программирование: ` + carModel.Programming.Name + `.
	//Код: ` + carModel.Code + `.
	//Пульт: ` + carModel.Pult.Name + `.`
	//Описание: ` + carModel.Description + `.`
}

func (c CarModelDetailScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c CarModelDetailScreen) GetScreenName() string {
	return "carModelDetail"
}
