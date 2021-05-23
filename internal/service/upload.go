package service

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *ServiceStruct) UploadResume(ctx *gin.Context) {
	file, err := ctx.FormFile("resume")
	if err != nil {
		log.Println(err)
	}

	accessDetails, _ := s.usecase.ExtractTokenMetadata(ctx.Request)

	path, err := filepath.Abs("file/resume")
	if err != nil {
		log.Println(err)
	}

	err = ctx.SaveUploadedFile(file, path+"/"+accessDetails.Username+".pdf")
	if err != nil {
		log.Println(err)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

}

func (s *ServiceStruct) UploadProfilePicture(ctx *gin.Context) {
	file, err := ctx.FormFile("profile picture")
	if err != nil {
		log.Println(err)
	}

	accessDetails, _ := s.usecase.ExtractTokenMetadata(ctx.Request)

	path, err := filepath.Abs("file/profilePicture")
	if err != nil {
		log.Println(err)
	}

	err = ctx.SaveUploadedFile(file, path+"/"+accessDetails.Username+".jpeg")
	if err != nil {
		log.Println(err)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

}
