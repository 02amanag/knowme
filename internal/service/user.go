package service

import (
	"net/http"

	"github.com/02amanag/p-02/internal/helper"
	"github.com/02amanag/p-02/internal/model"
	"github.com/gin-gonic/gin"
)

func (s *ServiceStruct) Login(ctx *gin.Context) {

	var entity model.Login

	if ctx.ShouldBindJSON(&entity) != nil {
		helper.DisplayError(ctx, "Invalid form", 406)
		return
	}

	token, username, err := s.usecase.Login(entity.Email, entity.Password)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"Message": "User signed in", "Token": token, "Username": username})
	} else {
		helper.DisplayError(ctx, err.Error(), 406)
	}
}

func (s ServiceStruct) SingUp(ctx *gin.Context) {
	var entity model.SingUp

	if ctx.ShouldBindJSON(&entity) != nil {
		helper.DisplayError(ctx, "Invalid form", 406)
		return
	}

	err := s.usecase.SingUp(entity.EmailId, entity.Password, entity.FirstName, entity.LastName)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"Message": "User signed up"})
	} else {
		helper.DisplayError(ctx, err.Error(), 406)
	}
}
