package service

import (
	"log"
	"path/filepath"

	"github.com/02amanag/p-02/internal/helper"
	"github.com/gin-gonic/gin"
)

func (s *ServiceStruct) GenerateQr(ctx *gin.Context) {
	accessDetails, _ := s.usecase.ExtractTokenMetadata(ctx.Request)
	err := s.usecase.GenerateQr(accessDetails.Username)
	if err != nil {
		helper.DisplayError(ctx, err.Error(), 404)
	} else {
		path, err := filepath.Abs("file/qr")
		if err != nil {
			log.Fatal(err)
		}
		ctx.FileAttachment(path+"/qr.png", accessDetails.Username)
	}
}

func (s *ServiceStruct) GetQr(ctx *gin.Context) {
	username := ctx.Param("username")
	err := s.usecase.GenerateQr(username)
	if err != nil {
		helper.DisplayError(ctx, err.Error(), 404)
	} else {
		path, err := filepath.Abs("file/qr")
		if err != nil {
			log.Fatal(err)
		}
		ctx.FileAttachment(path+"/qr.png", username+"_qr.png")
	}
}
