package DoorLockModel

import (
	"github.com/jinzhu/gorm"
	"log"
)

type DoorsLockModel struct {
	ID                   uint   `json:"id" gorm:"primary_key;column:id"`
	Name                 string `json:"name" gorm:"column:name"`
	DoorLockMarkId       uint64 `json:"doors_lock_mark_id" gorm:"column:doors_lock_mark_id"`
	LockTypeId           string `json:"lock_type_id" gorm:"column:lock_type_id"`
	LockMechSecretTypeId string `json:"lock_mech_secret_type_id" gorm:"column:lock_mech_secret_type_id"`
	SecretType           string `json:"secret_type" gorm:"column:secret_type"`
	ResistanceClass      string `json:"resistance_class" gorm:"column:resistance_class"`
	CreatedAt            string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt            string `json:"updated_at" gorm:"column:updated_at"`
}

type Storage struct {
	*gorm.DB
}

func (r *Storage) GetByMarkId(markId uint64) []DoorsLockModel {
	var doorsLockModels []DoorsLockModel

	result := r.Find(&doorsLockModels, "doors_lock_mark_id = ?", markId)

	if result.Error != nil {
		log.Fatalf("Ошибка при получении данных: %v", result.Error)
	}
	return doorsLockModels
}
