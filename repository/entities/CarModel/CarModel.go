package CarModel

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/ApiClientBackend"
)

type CarModel struct {
	ID                 uint64 `json:"id" gorm:"primary_key;column:id"`
	Name               string `json:"name" gorm:"column:name"`
	Description        string `json:"description"`
	Profile            string `json:"profile"`
	OpenInside         bool   `json:"open_inside"`
	OpenByClockArrowUp bool   `json:"open_by_clock_arrow_up"`
	Transponder        string `json:"transponder"`
	Prepare            string `json:"prepare"`
	Pult               string `json:"pult"`
	Programming        string `json:"programming"`
	Code               string `json:"code"`
}

//type Storage struct {
//	*gorm.DB
//}

func GetById(Id uint64) *CarModel {
	var carModel CarModel
	//
	//result := r.Preload("Prepare").Preload("Transponder").Preload("Pult").Preload("Programming").Preload("CarMark").First(&carModel, "id = ?", Id)
	//
	//if result.Error != nil {
	//	log.Println("Ошибка при получении данных: %v", result.Error)
	//	return nil
	//}
	return &carModel
}

func GetByMarkId(markId uint64, page uint, pageSize uint) []CarModel {

	url := fmt.Sprintf("car-models?limit=%d&page=%d&car_mark_id=%d", pageSize, page, markId)
	result := ApiClientBackend.Get(url)

	var carModels []CarModel

	for _, item := range result.Data {
		idFloat, _ := item["id"].(float64)
		carModel := CarModel{
			ID:          uint64(idFloat),
			Name:        item["name"].(string),
			Description: item["description"].(string),
			//Profile:            item["profile"].(string),
			OpenInside:         item["open_inside"].(bool),
			OpenByClockArrowUp: item["open_by_clock_arrow_up"].(bool),
			Transponder:        item["transponder"].(string),
			Prepare:            item["prepare"].(string),
			Pult:               item["pult"].(string),
			Programming:        item["programming"].(string),
			Code:               item["code"].(string),
		}
		carModels = append(carModels, carModel)

	}
	return carModels
}

//func (r *Storage) FindByName(name string) []CarModel {
//	var carModels []CarModel
//	r.Where("name ILIKE ?", "%"+name+"%").Find(&carModels).Order("name asc")
//
//	return carModels
//}
