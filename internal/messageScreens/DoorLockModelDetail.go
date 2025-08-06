package messageScreens

import (
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi/messageType"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/Attachment"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/DoorLockModel"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type DoorLockModelDetailScreen struct {
	OutputMessage messageType.OutputMessage
}

func NewDoorLockModelDetail(dto BaseScreenDTO) (baseScreen BaseScreen, clearPrevMessage bool) {
	var message messageType.OutputMessage
	s := DoorLockModel.Storage{DB: dto.DB}
	modelId, _ := strconv.ParseUint(dto.Filter["id"], 10, 32)
	doorsLockModel := s.GetById(modelId)
	if doorsLockModel == nil {
		return NewAlert(dto, "Не найдено.")
	}

	var text = getTextDoorModel(*doorsLockModel)

	a := Attachment.Storage{DB: dto.DB}
	attachmentables := a.GetForDoorLockModel(doorsLockModel.ID)

	if len(attachmentables) > 0 {
		message = messageWithAttachments(attachmentables, dto, text)
	} else {
		message = messageText(dto, text)
	}

	var baseScreenInterface BaseScreen = DoorLockModelDetailScreen{OutputMessage: message}

	return baseScreenInterface, true
}

func getTextDoorModel(doorsLockModel DoorLockModel.DoorsLockModel) string {

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

	return `Название: ` + doorsLockModel.Name + `.
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
}

func messageWithAttachments(attachmentables []Attachment.Attachmentable, dto BaseScreenDTO, text string) messageType.MessageWithImagesGroup {
	var mediaGroup []interface{}

	for i, attachmentable := range attachmentables {

		if attachmentable.Attachment.Alt != "" {
			media := tgbotapi.NewInputMediaPhoto(tgbotapi.FileID(attachmentable.Attachment.Alt))
			if i == 0 {
				media.Caption = text
			}
			mediaGroup = append(mediaGroup, media)
		}
	}

	return messageType.MessageWithImagesGroup{
		Bot:    dto.Bot,
		ChatId: dto.User.TgUserId,
		Media:  mediaGroup,
	}
}

func messageText(dto BaseScreenDTO, text string) messageType.TextMessage {

	return messageType.TextMessage{
		Text:   text,
		Bot:    dto.Bot,
		ChatId: dto.User.TgUserId,
	}
}

func (c DoorLockModelDetailScreen) GetOutputMessage() messageType.OutputMessage {
	return c.OutputMessage
}

func (c DoorLockModelDetailScreen) GetScreenName() string {
	return "doorLockModelsDetail"
}
