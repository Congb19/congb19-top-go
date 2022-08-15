package controller

import (
	"congb19-top-go/src/data"
	"congb19-top-go/src/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetClubs(context *gin.Context) {
	var request model.Clubs
	clubs := data.DB.Model(&model.Clubs{}).Find(&request)
	if clubs.Error != nil {
		panic(clubs.Error)
	}
	fmt.Println("clubs: ", clubs)
	if clubs.RowsAffected > 0 {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": clubs.Value})
	} else {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil})
	}
}
