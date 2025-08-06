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

func (r *Storage) GetAll(page uint, pageSize uint) []CarMark {
	var carMarks []CarMark
	offset := int(page) * int(pageSize)
	//result := r.Limit(pageSize).Offset(offset).Find(&carMarks).Order("name asc")

	sql := `
    SELECT DISTINCT car_marks.*
    FROM car_marks
    JOIN car_models ON car_models.car_mark_id = car_marks.id
    WHERE car_models.name IS NOT NULL 
         AND  car_models.description IS NOT NULL
         AND  car_models.profile IS NOT NULL
         AND  car_models.transponder_id IS NOT NULL
         AND  car_models.prepare_id IS NOT NULL
         AND  car_models.pult_id IS NOT NULL
         AND  car_models.programming_id IS NOT NULL
         AND  car_models.code IS NOT NULL
    LIMIT ? OFFSET ?
`

	result := r.Raw(sql, pageSize, offset).Scan(&carMarks)

	if result.Error != nil {
		log.Fatalf("Ошибка при получении данных: %v", result.Error)
	}
	return carMarks
}
