package postgres

import (
	"github.com/Vlad06013/unlockerTG.git/internal/domain/entity"
	"github.com/jinzhu/gorm"
)

type botStorage struct {
	*gorm.DB
}

func NewBotStorage(pg *gorm.DB) *botStorage {
	return &botStorage{pg}
}

func (r *botStorage) GetAll() []entity.Bot {
	var bots []entity.Bot
	r.Find(&bots)
	return bots
}
