package config

import (
	"github.com/joho/godotenv"
)

type Manager struct {
}

func NewManager() Manager {
	return Manager{}
}

func (Manager) Get(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	value, err := godotenv.Read()

	if err != nil {
		panic(err)
	}

	return value[key]
}
