package Attachment

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Attachment struct {
	ID   uint   `json:"id" gorm:"primary_key;column:id"`
	Name string `json:"name" gorm:"column:name"`
	Alt  string `json:"alt" gorm:"column:alt"`
}

type Attachmentable struct {
	ID                 uint   `json:"id" gorm:"primary_key;column:id"`
	AttachmentableType string `json:"attachmentable_type" gorm:"column:attachmentable_type"`
	AttachmentableId   string `json:"attachmentable_id" gorm:"column:attachmentable_id"`
	AttachmentId       string `json:"attachment_id" gorm:"column:attachment_id"`

	Attachment Attachment `gorm:"foreignKey:attachment_id"`
}

func (Attachmentable) TableName() string {
	return "attachmentable"
}

type Storage struct {
	*gorm.DB
}

func (r *Storage) GetForDoorLockModel(Id uint) []Attachmentable {
	//var attachments []Attachment
	var attachmentables []Attachmentable

	result := r.Preload("Attachment").Where("attachmentable_type = ? AND attachmentable_id = ?", "App\\Models\\DoorsLockModel", Id).Find(&attachmentables)

	if result.Error != nil {
		log.Fatalf("Ошибка при получении данных: %v", result.Error)
	}
	return attachmentables
}
