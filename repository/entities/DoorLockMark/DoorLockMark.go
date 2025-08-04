package DoorLockMark

import (
	"github.com/jinzhu/gorm"
	"log"
)

type DoorsLockMark struct {
	ID        uint   `json:"id" gorm:"primary_key;column:id"`
	Name      string `json:"name" gorm:"column:name"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

type Storage struct {
	*gorm.DB
}

func (r *Storage) GetAll() []DoorsLockMark {
	var doorsLockMarks []DoorsLockMark
	result := r.Find(&doorsLockMarks)

	if result.Error != nil {
		// Обработка ошибки
		log.Fatalf("Ошибка при получении данных: %v", result.Error)
	}
	return doorsLockMarks
}
