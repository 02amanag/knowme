package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigService interface {
	GetConfig(string) (string, error)
}

type Configure struct {
	path string
}

func (c Configure) GetConfig(key string) (string, error) {
	err := godotenv.Load(c.path)
	if err != nil {
		log.Println("Error loading .env file -> " + c.path)
	}
	value, ok := os.LookupEnv(key)
	if !ok {
		return value, errors.New("Error in fetching value from " + c.path)
	}
	return value, nil
}

func NewConfig(envFilePath string) ConfigService {
	return &Configure{
		path: envFilePath,
	}
}
