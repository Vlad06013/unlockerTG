package Setting

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Setting struct {
	ID          uint64 `json:"id" gorm:"primary_key;column:id"`
	SectionSlug string `json:"section_slug" gorm:"column:section_slug"`
	Name        string `json:"name" gorm:"column:name"`
	Slug        string `json:"slug" gorm:"column:slug"`
	Value       string `json:"value" gorm:"column:value"`
	Description string `json:"description" gorm:"column:description"`
	CreatedAt   string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   string `json:"updated_at" gorm:"column:updated_at"`
}

type Storage struct {
	*gorm.DB
}

func (r *Storage) GetBySlug(slug string) Setting {
	var setting Setting
	result := r.Find(&setting, "slug = ?", slug)

	if result.Error != nil {
		log.Fatalf("Ошибка при получении данных: %v", result.Error)
	}
	return setting
}
