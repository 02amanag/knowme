package service

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *ServiceStruct) DownloadResume(ctx *gin.Context) {
	username := ctx.Param("username")
	path, err := filepath.Abs("file/resume")
	if err != nil {
		log.Fatal(err)
	}

	ctx.FileAttachment(path+"/"+username+".pdf", username+"_resume.pdf")
}
