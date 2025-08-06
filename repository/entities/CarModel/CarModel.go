package CarModel

import (
	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarMark"
	"github.com/jinzhu/gorm"
	"log"
)

type CarModel struct {
	ID                 uint64 `json:"id" gorm:"primary_key;column:id"`
	Name               string `json:"name" gorm:"column:name"`
	CarMarkId          string `json:"car_mark_id" gorm:"column:car_mark_id"`
	Description        string `json:"description" gorm:"column:description"`
	Profile            string `json:"profile" gorm:"column:profile"`
	OpenInside         bool   `json:"open_inside" gorm:"column:open_inside"`
	OpenByClockArrowUp bool   `json:"open_by_clock_arrow_up" gorm:"column:open_by_clock_arrow_up"`
	TransponderId      uint64 `json:"transponder_id" gorm:"column:transponder_id"`
	PrepareId          uint64 `json:"prepare_id" gorm:"column:prepare_id"`
	PultId             uint64 `json:"pult_id" gorm:"column:pult_id"`
	ProgrammingId      uint64 `json:"programming_id" gorm:"column:programming_id"`
	Code               string `json:"code" gorm:"column:code"`
	CreatedAt          string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt          string `json:"updated_at" gorm:"column:updated_at"`

	CarMark     CarMark.CarMark `json:"car_mark" gorm:"foreignKey:CarMarkId"`
	Prepare     Prepare
	Transponder Transponder
	Pult        Pult
	Programming Programming
}

type Storage struct {
	*gorm.DB
}

type Prepare struct {
	ID   uint64 `json:"id" gorm:"primary_key;column:id"`
	Name string `json:"name" gorm:"column:name"`
}

type Transponder struct {
	ID   uint64 `json:"id" gorm:"primary_key;column:id"`
	Name string `json:"name" gorm:"column:name"`
}

type Pult struct {
	ID   uint64 `json:"id" gorm:"primary_key;column:id"`
	Name string `json:"name" gorm:"column:name"`
}

type Programming struct {
	ID   uint64 `json:"id" gorm:"primary_key;column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func (r *Storage) GetById(Id uint64) *CarModel {
	var carModel CarModel

	result := r.Preload("Prepare").Preload("Transponder").Preload("Pult").Preload("Programming").Preload("CarMark").First(&carModel, "id = ?", Id)

	if result.Error != nil {
		log.Println("Ошибка при получении данных: %v", result.Error)
		return nil
	}
	return &carModel
}

func (r *Storage) GetByMarkId(markId uint64, page uint, pageSize uint) []CarModel {
	var carModels []CarModel
	offset := int(page) * int(pageSize)

	//result := r.Limit(pageSize).Offset(offset).Find(&carModels, "car_mark_id = ?", markId)
	result := r.
		Limit(pageSize).
		Offset(offset).
		Where("car_mark_id = ?", markId).
		Where("name IS NOT NULL").
		Where("description IS NOT NULL").
		Where("profile IS NOT NULL").
		Where("transponder_id IS NOT NULL").
		Where("prepare_id IS NOT NULL").
		Where("pult_id IS NOT NULL").
		Where("programming_id IS NOT NULL").
		Where("code IS NOT NULL").
		Find(&carModels)

	if result.Error != nil {
		log.Fatalf("Ошибка при получении данных: %v", result.Error)
	}
	return carModels
}

func (r *Storage) FindByName(name string) []CarModel {
	var carModels []CarModel
	r.Where("name ILIKE ?", "%"+name+"%").Find(&carModels).Order("name asc")

	return carModels
}
