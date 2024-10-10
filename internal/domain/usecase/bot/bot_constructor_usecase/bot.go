package bot_constructor_usecase

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/adapters/repository/db/postgres"
	"github.com/Vlad06013/unlockerTG.git/internal/controller/api/telegram"
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
	"github.com/Vlad06013/unlockerTG.git/internal/domain/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jinzhu/gorm"
	"sync"
)

type BotUseCase struct {
	botService       service.BotService
	userService      service.UserService
	messageService   service.MessageService
	output           *telegram.Output
	messageConstruct entity.Constructable
}

func NewBotUseCase(db *gorm.DB) *BotUseCase {
	botStorage := postgres.NewBotStorage(db)
	botService := service.NewBotService(botStorage)

	userStorage := postgres.NewUserStorage(db)
	userService := service.NewUserService(userStorage)

	messageStorage := postgres.NewMessageStorage(db)
	messageService := service.NewMessageService(messageStorage)

	return &BotUseCase{*botService, *userService, *messageService, nil, nil}
}

func (bu BotUseCase) StartListenerUpdates() {
	bots := bu.botService.GetAll()
	if len(bots) == 0 {
		panic("Ботов нет")
	}
	var counts = len(bots)
	var botApi = make([]*entity.BotApi, counts)

	for i, bot := range bots {
		api, _ := tgbotapi.NewBotAPI(bot.Token)
		botApi[i] = &entity.BotApi{
			Api: *api,
			Bot: &bot,
		}
	}
	var wg sync.WaitGroup
	wg.Add(counts)

	work := func(bot *entity.BotApi) {
		defer wg.Done()
		fmt.Print("Listening upd for bot " + bot.Bot.Name)
		bu.ListenUpdates(bot)
	}
	for _, bot := range botApi {
		go work(bot)
	}
	wg.Wait()
}

func (bu BotUseCase) SaveLastMessage(messageId uint, messageTGId uint) {

	historyId := bu.userService.User.BotHistory.ID
	var newHistory = entity.TgUserMessageHistory{
		LastMessageId:   messageId,
		LastTGMessageId: messageTGId,
	}
	bu.userService.User.BotHistory = bu.userService.UpdateHistory(historyId, &newHistory)
}

func (bu BotUseCase) CheckSendNextMessage() bool {
	lastMessage := bu.userService.User.BotHistory.LastMessage
	if len(lastMessage.Keyboard.Buttons) == 0 && lastMessage.NextMessageId != 0 && lastMessage.Keyboard.TableName == "" {
		return true
	}
	return false
}
