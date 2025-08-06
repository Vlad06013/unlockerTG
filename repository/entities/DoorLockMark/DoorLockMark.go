package DoorLockMark

import (
	"github.com/jinzhu/gorm"
	"log"
)

type DoorsLockMark struct {
	ID        uint64 `json:"id" gorm:"primary_key;column:id"`
	Name      string `json:"name" gorm:"column:name"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at"`
}

type Storage struct {
	*gorm.DB
}

func (r *Storage) GetAll(page uint, pageSize uint) []DoorsLockMark {
	offset := int(page) * int(pageSize)

	var doorsLockMarks []DoorsLockMark
	//result := r.Limit(pageSize).Offset(offset).Find(&doorsLockMarks)

	sql := `
	  SELECT DISTINCT doors_lock_marks.*
	  FROM doors_lock_marks
	  JOIN doors_lock_models ON doors_lock_models.doors_lock_mark_id = doors_lock_marks.id
	  WHERE doors_lock_models.name IS NOT NULL
	       AND  doors_lock_models.lock_type_id IS NOT NULL
	       AND  doors_lock_models.lock_mech_secret_type_id IS NOT NULL
	       AND  doors_lock_models.secret_type IS NOT NULL
	       AND  doors_lock_models.resistance_class IS NOT NULL
	       AND  doors_lock_models.center_distance IS NOT NULL
	       AND  doors_lock_models.backset IS NOT NULL
	       AND  doors_lock_models.end_strip_length IS NOT NULL
	       AND  doors_lock_models.center_distance_fastenings IS NOT NULL
	       AND  doors_lock_models.crossbar_diameter IS NOT NULL
	       AND  doors_lock_models.deadbolt_overhang IS NOT NULL
	       AND  doors_lock_models.overhang_count IS NOT NULL
	       AND  doors_lock_models.body_height IS NOT NULL
	       AND  doors_lock_models.case_depth IS NOT NULL
	       AND  doors_lock_models.width_depth IS NOT NULL
	       AND  doors_lock_models.locking_from_inside IS NOT NULL
	       AND  doors_lock_models.key_type IS NOT NULL
	       AND  doors_lock_models.description IS NOT NULL
		ORDER BY doors_lock_marks.name ASC
	  LIMIT ? OFFSET ?
	`

	result := r.Raw(sql, pageSize, offset).Scan(&doorsLockMarks)

	if result.Error != nil {
		// Обработка ошибки
		log.Fatalf("Ошибка при получении данных: %v", result.Error)
	}
	return doorsLockMarks
}
