package main

import (
	"github.com/Vlad06013/unlockerTG.git/config"
	"github.com/Vlad06013/unlockerTG.git/infrastructure/tgBotApi"
	dbClient "github.com/Vlad06013/unlockerTG.git/repository"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jinzhu/gorm"
)

func main() {
	conn := initDB()
	tgBotApi.Listen(conn)
}

func initDB() *gorm.DB {
	err := config.SetEnvValues()
	if err != nil {
		panic(err)
	}
	var cfg config.DBPostgres
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}

	return dbClient.NewConnection(cfg)
}
