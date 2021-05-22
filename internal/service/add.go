package service

import (
	"net/http"

	"github.com/02amanag/p-02/internal/helper"
	"github.com/02amanag/p-02/internal/model"
	"github.com/gin-gonic/gin"
)

func (s *ServiceStruct) AddProfile(ctx *gin.Context) {
	var entity model.AddData
	if ctx.ShouldBindJSON(&entity) != nil {
		helper.DisplayError(ctx, "Invalid form", 406)
		return
	}
	accessDetails, _ := s.usecase.ExtractTokenMetadata(ctx.Request)
	err := s.usecase.AddData(entity, accessDetails.Username)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Data Inserted"})
	} else {
		helper.DisplayError(ctx, err.Error(), 406)
	}
}
