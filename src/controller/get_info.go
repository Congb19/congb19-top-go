package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHappinessList(context *gin.Context) {
	result:=123
	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": result})
}