package config

import (
	"fmt"
	"os"
)

type DBPostgres struct {
	DbHost     string `env:"DbHost"`
	DbPort     string `env:"DbPort"`
	DbName     string `env:"DbName"`
	DbUsername string `env:"DbUsername"`
	DbPassword string `env:"DbPassword"`
}

func SetEnvValues() error {
	err := os.Setenv("DbHost", "localhost")
	if err != nil {
		return fmt.Errorf("error setting port, err = %v", err)
	}
	err = os.Setenv("DbPort", "5432")
	if err != nil {
		return fmt.Errorf("error setting port, err = %v", err)
	}
	err = os.Setenv("DbName", "ulocker_lab")
	if err != nil {
		return fmt.Errorf("error setting port, err = %v", err)
	}

	err = os.Setenv("DbUsername", "ulocker_lab")
	if err != nil {
		return fmt.Errorf("error setting port, err = %v", err)
	}

	err = os.Setenv("DbPassword", "eRdHs3gFsa3gt!n")
	if err != nil {
		return fmt.Errorf("error setting port, err = %v", err)
	}

	return nil
}
