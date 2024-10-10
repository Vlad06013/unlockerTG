package main

import (
	"github.com/Vlad06013/unlockerTG.git/internal/config"
	"github.com/Vlad06013/unlockerTG.git/internal/domain/usecase/bot/bot_constructor_usecase"
	postgres2 "github.com/Vlad06013/unlockerTG.git/pkg/client/postgres"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	err := config.SetEnvValues()
	if err != nil {
		panic(err)
	}

	var cfg config.ConfigDBPostgres
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}
	var conn = postgres2.NewConnection(cfg)
	bu := bot_constructor_usecase.NewBotUseCase(conn)
	bu.StartListenerUpdates()
}
