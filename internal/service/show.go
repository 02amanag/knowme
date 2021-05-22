package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *ServiceStruct) GetProfile(ctx *gin.Context) {
	username := ctx.Param("username")

	response := s.usecase.GetProfile(username)
	ctx.HTML(http.StatusOK, "index.html", response)
	// ctx.JSON(http.StatusOK, gin.H{"Message": "User signed in", "response": response})
}
