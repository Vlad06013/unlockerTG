package service

import (
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
)

type UserStorage interface {
	GetUserByTGID(tgID int64) (*entity.TgUser, error)
	CreateUser(u entity.TgUser) *entity.TgUser
	CreateUserHistory(TgUserId uint, botId uint) *entity.TgUserMessageHistory
	GetHistoryByBotId(tgUserId int64, botId uint) (entity.TgUserMessageHistory, error)
	UpdateHistory(Id uint, history *entity.TgUserMessageHistory) *entity.TgUserMessageHistory
}

type UserService struct {
	storage UserStorage
	User    *entity.TgUser
}

func NewUserService(storage UserStorage) *UserService {
	return &UserService{storage: storage}
}

func (s UserService) GetUserByTGID(tgID int64) (*entity.TgUser, error) {
	user, err := s.storage.GetUserByTGID(tgID)
	return user, err
}

func (s *UserService) CreateUser(u entity.TgUser) *entity.TgUser {
	user := s.storage.CreateUser(u)
	return user
}

func (s *UserService) GetHistoryByBotId(tgUserId int64, botId uint) (entity.TgUserMessageHistory, error) {
	history, err := s.storage.GetHistoryByBotId(tgUserId, botId)
	return history, err
}

func (s *UserService) CreateUserHistory(TgUserId uint, botId uint) *entity.TgUserMessageHistory {

	history := s.storage.CreateUserHistory(TgUserId, botId)
	return history
}

func (s *UserService) UpdateHistory(Id uint, history *entity.TgUserMessageHistory) *entity.TgUserMessageHistory {
	return s.storage.UpdateHistory(Id, history)

}

func (s *UserService) InitUser(tgID int64, name string, bot *entity.Bot) *entity.TgUser {

	user, err := s.GetUserByTGID(tgID)
	if err != nil {
		var userEntity = entity.TgUser{
			TgUserId: tgID,
			Name:     name,
		}
		user = s.CreateUser(userEntity)
		user.BotHistory = s.CreateUserHistory(user.ID, bot.ID)
	}
	if user.BotHistory == nil {
		user.BotHistory = s.CreateUserHistory(user.ID, bot.ID)
	}
	s.User = user
	return user
}
