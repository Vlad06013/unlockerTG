package messageScreens

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"reflect"
)

type ButtonConstructorDTO struct {
	FieldForText     string
	PrefixCallback   string
	FieldForCallback string
	Entities         interface{}
}

var keyboard tgbotapi.InlineKeyboardMarkup

func GenerateButtons(dto ButtonConstructorDTO) [][]tgbotapi.InlineKeyboardButton {
	var row []tgbotapi.InlineKeyboardButton
	var rows [][]tgbotapi.InlineKeyboardButton

	slice := reflect.ValueOf(dto.Entities)
	// Объявляем переменную для индекса ряда
	var rowCountIndex int

	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)

		// Если элемент указатель, получаем его значение
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		// Проверяем, что это структура
		if item.Kind() != reflect.Struct {
			continue // или обработка ошибки
		}

		// Получаем поле по имени из dto
		btnTextVal := item.FieldByName(dto.FieldForText)
		cbValue := item.FieldByName(dto.FieldForCallback)

		btnText := fmt.Sprintf("%v", btnTextVal.Interface())
		cbFilter := fmt.Sprintf("%v", cbValue.Interface())

		callbackData := dto.PrefixCallback + "_" + cbFilter

		button := tgbotapi.NewInlineKeyboardButtonData(btnText, callbackData)
		row = append(row, button)

		currentCountInRow := countInRowOptions[rowCountIndex]

		if (len(row) == currentCountInRow) || i == slice.Len()-1 {
			rows = append(rows, row)
			row = nil
			rowCountIndex = (rowCountIndex + 1) % len(countInRowOptions)
		}
	}
	return rows
}
