package CarMark

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/internal/ApiClientBackend"
	"github.com/jinzhu/gorm"
)

type CarMark struct {
	ID   uint64 `json:"id" gorm:"primary_key;column:id"`
	Name string `json:"name" gorm:"column:name"`
	//CreatedAt string `json:"created_at" gorm:"column:created_at"`
	//UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

type Storage struct {
	*gorm.DB
}

func (r *Storage) GetAll(page uint, pageSize uint) []CarMark {
	//var carMarks []CarMark
	//offset := int(page) * int(pageSize)
	////result := r.Limit(pageSize).Offset(offset).Find(&carMarks).Order("name asc")
	//
	//sql := `
	//  SELECT DISTINCT car_marks.*
	//  FROM car_marks
	//  JOIN car_models ON car_models.car_mark_id = car_marks.id
	//  WHERE car_models.name IS NOT NULL
	//       AND  car_models.description IS NOT NULL
	//       AND  car_models.profile IS NOT NULL
	//       AND  car_models.transponder_id IS NOT NULL
	//       AND  car_models.prepare_id IS NOT NULL
	//       AND  car_models.pult_id IS NOT NULL
	//       AND  car_models.programming_id IS NOT NULL
	//       AND  car_models.code IS NOT NULL
	//  ORDER BY car_marks.name ASC
	//  LIMIT ? OFFSET ?
	//`
	//
	//result := r.Raw(sql, pageSize, offset).Scan(&carMarks)
	//
	//if result.Error != nil {
	//	log.Fatalf("Ошибка при получении данных: %v", result.Error)
	//}
	//return carMarks
	return getAll(page, pageSize)

}

func getAll(page uint, pageSize uint) []CarMark {

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

func (r *Storage) FindByName(name string) []CarMark {
	var carMarks []CarMark
	r.Where("name ILIKE ?", "%"+name+"%").Find(&carMarks).Order("name asc")

	return carMarks
}
