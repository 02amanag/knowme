package service

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *ServiceStruct) UploadResume(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	accessDetails, _ := s.usecase.ExtractTokenMetadata(ctx.Request)

	path, err := filepath.Abs("file/resume")
	if err != nil {
		log.Fatal(err)
	}

	err = ctx.SaveUploadedFile(file, path+"/"+accessDetails.Username)
	if err != nil {
		log.Fatal(err)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

}

func (s *ServiceStruct) UploadProfilePicture(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	accessDetails, _ := s.usecase.ExtractTokenMetadata(ctx.Request)

	path, err := filepath.Abs("file/profilePicture")
	if err != nil {
		log.Fatal(err)
	}

	err = ctx.SaveUploadedFile(file, path+"/"+accessDetails.Username)
	if err != nil {
		log.Fatal(err)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

}