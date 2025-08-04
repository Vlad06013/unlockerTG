package repository

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewConnection(dbc config.DBPostgres) *gorm.DB {

	db, err := gorm.Open("postgres", "host="+dbc.DbHost+" user="+dbc.DbUsername+" password="+dbc.DbPassword+" dbname="+dbc.DbName+" port="+dbc.DbPort+" sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("Не удалось подключиться к базе данных")
	}
	return db
}
