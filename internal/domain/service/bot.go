package service

import (
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
)

type BotStorage interface {
	GetAll() []entity.Bot
}

type BotService struct {
	storage BotStorage
}

func NewBotService(storage BotStorage) *BotService {
	return &BotService{
		storage: storage,
	}
}

func (s BotService) GetAll() []entity.Bot {
	return s.storage.GetAll()
}
