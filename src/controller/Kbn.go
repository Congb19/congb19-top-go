package controller

import (
	"congb19-top-go/src/data"
	"congb19-top-go/src/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHappinessList(context *gin.Context) {
	var items []model.KbnItems
	result := data.DB.Where(&model.KbnItems{Type: 1}).Where(&model.KbnItems{Show: true}).Find(&items)
	if result.Error != nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "ServerError", "data": nil})
		panic(result.Error)
	}
	fmt.Println("items: ", items)
	if result.RowsAffected > 0 {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Found", "data": items})
	} else {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "RecordNotFound", "data": nil})
	}
}
func PostKbn(context *gin.Context) {
	var kbn model.KbnItems
	if err := context.ShouldBindJSON(&kbn); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		context.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error(), "data": nil})
		return
	}
	fmt.Println("get kbn: ", &kbn)

	kbn.Show = false

	result := data.DB.Create(&kbn)
	if result.Error != nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "ServerError", "data": nil})
		panic(result.Error)
	}

	context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "share ok", "data": nil})
}
