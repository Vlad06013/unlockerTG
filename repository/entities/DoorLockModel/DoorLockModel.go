package DoorLockModel

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/DoorLockMark"
	"github.com/jinzhu/gorm"
	"log"
)

type DoorsLockModel struct {
	ID                       uint   `json:"id" gorm:"primary_key;column:id"`
	Name                     string `json:"name" gorm:"column:name"`
	DoorLockMarkId           uint64 `json:"doors_lock_mark_id" gorm:"column:doors_lock_mark_id"`
	LockTypeId               uint64 `json:"lock_type_id" gorm:"column:lock_type_id"`
	LockMechSecretTypeId     uint64 `json:"lock_mech_secret_type_id" gorm:"column:lock_mech_secret_type_id"`
	SecretType               string `json:"secret_type" gorm:"column:secret_type"`
	ResistanceClass          uint64 `json:"resistance_class" gorm:"column:resistance_class"`
	TailLatch                bool   `json:"tail_latch" gorm:"column:tail_latch"`
	LatchInside              bool   `json:"latch_inside" gorm:"column:latch_inside"`
	Rods                     bool   `json:"rods" gorm:"column:rods"`
	CenterDistance           string `json:"center_distance" gorm:"column:center_distance"`
	Backset                  string `json:"backset" gorm:"column:backset"`
	EndStripLength           string `json:"end_strip_length" gorm:"column:end_strip_length"`
	EndStripWidth            string `json:"end_strip_width" gorm:"column:end_strip_width"`
	CenterDistanceFastenings string `json:"center_distance_fastenings" gorm:"column:center_distance_fastenings"`
	CrossbarDiameter         string `json:"crossbar_diameter" gorm:"column:crossbar_diameter"`
	DeadboltOverhang         string `json:"deadbolt_overhang" gorm:"column:deadbolt_overhang"`
	OverhangCount            uint64 `json:"overhang_count" gorm:"column:overhang_count"`
	BodyHeight               string `json:"body_height" gorm:"column:body_height"`
	CaseDepth                string `json:"case_depth" gorm:"column:case_depth"`
	WidthDepth               string `json:"width_depth" gorm:"column:width_depth"`
	LockingFromInside        string `json:"locking_from_inside" gorm:"column:locking_from_inside"`
	KeyType                  string `json:"key_type" gorm:"column:key_type"`
	Description              string `json:"description" gorm:"column:description"`
	CreatedAt                string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt                string `json:"updated_at" gorm:"column:updated_at"`

	LockType           LockType
	LockMechSecretType LockMechSecretType
	DoorLockMark       DoorLockMark.DoorsLockMark `gorm:"foreignKey:DoorLockMarkId"`
}

type LockType struct {
	ID   uint   `json:"id" gorm:"primary_key;column:id"`
	Name string `json:"name" gorm:"column:name"`
}

type LockMechSecretType struct {
	ID   uint   `json:"id" gorm:"primary_key;column:id"`
	Name string `json:"name" gorm:"column:name"`
}

type Storage struct {
	*gorm.DB
}

func (r *Storage) GetById(Id uint64) *DoorsLockModel {
	var doorsLockModels DoorsLockModel

	result := r.Preload("LockType").Preload("LockMechSecretType").Preload("DoorLockMark").First(&doorsLockModels, "id = ?", Id)

	if result.Error != nil {
		log.Println("Ошибка при получении данных: %v", result.Error)
		return nil
	}
	return &doorsLockModels
}

func (r *Storage) GetByMarkId(markId uint64) []DoorsLockModel {
	var doorsLockModels []DoorsLockModel
	fmt.Print(markId)
	result := r.Find(&doorsLockModels, "doors_lock_mark_id = ?", markId)

	if result.Error != nil {
		log.Println("Ошибка при получении данных: %v", result.Error)
	}
	return doorsLockModels
}
