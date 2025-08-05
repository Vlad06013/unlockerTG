package CarMark

import (
	"github.com/jinzhu/gorm"
	"log"
)

type CarMark struct {
	ID        uint64 `json:"id" gorm:"primary_key;column:id"`
	Name      string `json:"name" gorm:"column:name"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

type Storage struct {
	*gorm.DB
}

func (r *Storage) GetAll() []CarMark {
	var carMarks []CarMark
	result := r.Find(&carMarks)

	if result.Error != nil {
		// Обработка ошибки
		log.Fatalf("Ошибка при получении данных: %v", result.Error)
	}
	return carMarks
}
