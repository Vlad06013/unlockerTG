package messageScreens

import (
	//"fmt"
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/DoorLockModel"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/TgUser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"strconv"
)

type DoorLockModelDetailScreen struct {
	OutputMessage messageType.OutputMessage
}

func DoorLockModelDetail(user TgUser.TgUser, bot tgbotapi.BotAPI, db *gorm.DB, filter *uint64) BaseScreen {
	s := DoorLockModel.Storage{DB: db}
	doorsLockModel := s.GetById(*filter)

	iconFalse := "❌"
	iconTrue := "✅"
	tailLatch, latchInside, rods := iconFalse, iconFalse, iconFalse

	if doorsLockModel.TailLatch {
		tailLatch = iconTrue
	}

	if doorsLockModel.LatchInside {
		latchInside = iconTrue
	}

	if doorsLockModel.Rods {
		rods = iconTrue
	}

	var text = `Название: ` + doorsLockModel.Name + `.
	Производитель: ` + doorsLockModel.DoorLockMark.Name + `.
	Тип замка: ` + doorsLockModel.LockType.Name + `.
	Тип механизма секретности: ` + doorsLockModel.LockMechSecretType.Name + `.
	Тип секретности: ` + doorsLockModel.SecretType + `.
	Класс Взломостойкости: ` + strconv.FormatUint(doorsLockModel.ResistanceClass, 10) + `.
	Фалевая защелка: ` + tailLatch + `.
	Защелка изнутри: ` + latchInside + `.
	Наличие тяг: ` + rods + `.
	Межосевое расстояние - мм: ` + doorsLockModel.CenterDistance + `.
	Бэксет (удаление ключевого отверстия) - мм: ` + doorsLockModel.Backset + `.
	Длина торцевой планки - мм: ` + doorsLockModel.EndStripLength + `.
	Ширина торцевой планки - мм: ` + doorsLockModel.EndStripWidth + `.
	Межосевое расстояние креплений замка - мм: ` + doorsLockModel.CenterDistanceFastenings + `.
	Диаметр ригеля(Высота если квадратный) - мм: ` + doorsLockModel.CrossbarDiameter + `.
	Вылет ригеля - мм: ` + doorsLockModel.DeadboltOverhang + `.
	Кол-во ригелей: ` + strconv.FormatUint(doorsLockModel.OverhangCount, 10) + `.
	Высота корпуса замка - мм: ` + doorsLockModel.BodyHeight + `.
	Глубина корпуса замка - мм: ` + doorsLockModel.CaseDepth + `.
	Ширина корпуса замка - мм: ` + doorsLockModel.WidthDepth + `.
	Тип ключа: ` + doorsLockModel.KeyType + `.
	Запирание изнутри: ` + doorsLockModel.LockingFromInside + `.
	Описание: ` + doorsLockModel.Description + `.`

	media1 := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL("https://media.istockphoto.com/id/1325997570/ru/%D1%84%D0%BE%D1%82%D0%BE/%D0%B1%D0%B5%D0%BD%D0%B3%D0%B0%D0%BB%D1%8C%D1%81%D0%BA%D0%B0%D1%8F-%D0%BA%D0%BE%D1%88%D0%BA%D0%B0-%D0%BB%D0%B5%D0%B6%D0%B8%D1%82-%D0%BD%D0%B0-%D0%B4%D0%B8%D0%B2%D0%B0%D0%BD%D0%B5-%D0%B8-%D1%83%D0%BB%D1%8B%D0%B1%D0%B0%D0%B5%D1%82%D1%81%D1%8F.jpg?s=1024x1024&w=is&k=20&c=8s86dGhsMc42C0a7eDDkO0eSATEzMkQApbW3M0BxYG8="))
	media2 := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL("https://media.istockphoto.com/id/1325997570/ru/%D1%84%D0%BE%D1%82%D0%BE/%D0%B1%D0%B5%D0%BD%D0%B3%D0%B0%D0%BB%D1%8C%D1%81%D0%BA%D0%B0%D1%8F-%D0%BA%D0%BE%D1%88%D0%BA%D0%B0-%D0%BB%D0%B5%D0%B6%D0%B8%D1%82-%D0%BD%D0%B0-%D0%B4%D0%B8%D0%B2%D0%B0%D0%BD%D0%B5-%D0%B8-%D1%83%D0%BB%D1%8B%D0%B1%D0%B0%D0%B5%D1%82%D1%81%D1%8F.jpg?s=1024x1024&w=is&k=20&c=8s86dGhsMc42C0a7eDDkO0eSATEzMkQApbW3M0BxYG8="))
	media1.Caption = text
	mediaGroup := []interface{}{media1, media2}

	var message = messageType.MessageWithImagesGroup{
		Bot:    bot,
		ChatId: user.TgUserId,
		Media:  mediaGroup,
	}

	var screen = DoorLockMarksScreen{
		OutputMessage: messageType.OutputMessage(message),
	}

	var baseScreen BaseScreen = screen

	return baseScreen
}

func (c DoorLockModelDetailScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c DoorLockModelDetailScreen) GetScreenName() string {
	return "doorLockModelsDetail"
}
