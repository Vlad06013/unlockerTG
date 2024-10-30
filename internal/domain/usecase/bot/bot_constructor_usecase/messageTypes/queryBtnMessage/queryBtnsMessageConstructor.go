package queryBtnMessage

import (
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
	"github.com/Vlad06013/unlockerTG.git/internal/domain/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

type BtnQueryMessageConstructor struct {
	Message        *entity.Message
	ChatId         int64
	Text           string
	Type           string
	Keyboard       entity.Keyboard
	Buttons        []tgbotapi.InlineKeyboardButton
	MessageService *service.MessageService
}

func (c *BtnQueryMessageConstructor) TypeMessage() string {
	return c.Type
}

func (c *BtnQueryMessageConstructor) TextMessage() *string {
	return &c.Text
}

func (c *BtnQueryMessageConstructor) Chat() int64 {
	return c.ChatId
}

func (c *BtnQueryMessageConstructor) ButtonsMessage() []tgbotapi.InlineKeyboardButton {
	return c.Buttons
}

func (c *BtnQueryMessageConstructor) GetMessage() *entity.Message {
	return c.Message
}

func getDataPrefix(message *entity.Message) string {
	var callbackDataQuery, prefix string
	if message.NextMessageId != 0 {
		prefix = "mess"
		messageId := strconv.FormatUint(uint64(message.NextMessageId), 10)
		callbackDataQuery = prefix + "_" + messageId + "/filter_"

	} else {
		prefix = "query"
		callbackDataQuery = prefix + "_"

	}
	return callbackDataQuery
}

func (c *BtnQueryMessageConstructor) generateButtons() []tgbotapi.InlineKeyboardButton {
	var buttons []tgbotapi.InlineKeyboardButton
	var buttonText, callbackDataValue, callbackData string

	queryButtons := c.Keyboard.QueryButtons

	if queryButtons != nil {
		callbackDataQuery := getDataPrefix(c.Message)

		for queryButtons.Next() {
			queryButtons.Scan(&buttonText, &callbackDataValue)

			callbackData = callbackDataQuery + callbackDataValue

			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackData))

		}
	}
	callBackData := c.checkBackBtn()
	if callBackData != "" {
		buttons = addBackBtn(buttons, callBackData)
	}
	return buttons
}

func (c *BtnQueryMessageConstructor) checkBackBtn() string {
	if c.Message != nil {
		lastMessage, _ := c.MessageService.GetMessageWithFilter("next_message_id", c.Message.ID)
		if lastMessage != nil && (len(lastMessage.Keyboard.Buttons) != 0 || lastMessage.Keyboard.TableName != "") {
			return strconv.FormatUint(uint64(lastMessage.ID), 10)
		} else {
			messagable := c.MessageService.GetMessagableByNextMessage(c.Message.ID)
			if messagable != nil {
				return strconv.FormatUint(uint64(messagable.FromMessageId), 10)
			}
		}
	}
	return ""
}

func addBackBtn(buttons []tgbotapi.InlineKeyboardButton, callback string) []tgbotapi.InlineKeyboardButton {
	return addCustomBtn(buttons, "Назад", "mess_"+callback)
}

func addCustomBtn(buttons []tgbotapi.InlineKeyboardButton, text string, callback string) []tgbotapi.InlineKeyboardButton {
	return append(buttons, tgbotapi.NewInlineKeyboardButtonData(text, callback))
}

func NewMessage(construct BtnQueryMessageConstructor) entity.Constructable {

	buttons := construct.generateButtons()
	var queryBtnMessage entity.Constructable = &BtnQueryMessageConstructor{
		Message:  construct.Message,
		ChatId:   construct.ChatId,
		Text:     construct.Text,
		Type:     construct.Type,
		Keyboard: construct.Keyboard,
		Buttons:  buttons,
	}

	return queryBtnMessage
}
