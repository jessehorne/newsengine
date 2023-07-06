package util

import (
	"github.com/joho/godotenv"
)

func InitConf() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	return nil
}
