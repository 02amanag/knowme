package usecase

import (
	"github.com/02amanag/p-02/config"
	"github.com/02amanag/p-02/internal/repository"
)

type UsecaseStruct struct {
	config config.ConfigService
	repo   repository.RepositoryStruct
}

func NewUsecaseStruct(repo repository.RepositoryStruct, config config.ConfigService) *UsecaseStruct {
	return &UsecaseStruct{repo: repo, config: config}
}

func (u *UsecaseStruct) Healthy(param string) string {
	param += "added from usecase"
	return param
}
