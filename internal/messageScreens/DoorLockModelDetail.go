package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/Attachment"
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

	var mediaGroup []interface{}

	a := Attachment.Storage{DB: db}
	attachmentables := a.GetForDoorLockModel(doorsLockModel.ID)

	for i, attachmentable := range attachmentables {

		if attachmentable.Attachment.Alt != "" {
			media := tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(attachmentable.Attachment.Alt))
			if i == 0 {
				media.Caption = text
			}
			mediaGroup = append(mediaGroup, media)
		}
	}

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
