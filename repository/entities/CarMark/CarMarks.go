package CarMark

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/ApiClientBackend"
)

type CarMark struct {
	ID   uint64 `json:"id" gorm:"primary_key;column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func GetAll(page uint, pageSize uint) []CarMark {
	url := fmt.Sprintf("car-marks?limit=%d&page=%d", pageSize, page)
	result := ApiClientBackend.Get(url)

	var carMarks []CarMark

	for _, item := range result.Data {
		idFloat, _ := item["id"].(float64)
		carMark := CarMark{
			ID:   uint64(idFloat),
			Name: item["name"].(string),
		}
		carMarks = append(carMarks, carMark)

	}
	return carMarks
}

//func (r *Storage) FindByName(name string) []CarMark {
//	var carMarks []CarMark
//	r.Where("name ILIKE ?", "%"+name+"%").Find(&carMarks).Order("name asc")
//
//	return carMarks
//}
