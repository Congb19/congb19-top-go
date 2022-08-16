package controller

import (
	"congb19-top-go/src/data"
	"congb19-top-go/src/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetClubs(context *gin.Context) {
	var clubs []model.Clubs
	result := data.DB.Where(&model.Clubs{ValidFlag: true}).Find(&clubs)
	if result.Error != nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "ServerError", "data": nil})
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Found", "data": clubs})
	} else {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "RecordNotFound", "data": nil})
	}
}
func GetTools(context *gin.Context) {
	var tools []model.Tools
	id := context.DefaultQuery("id", "1")
	result := data.DB.Where("valid_flag = 1 and club_id = ?", id).Find(&tools)
	if result.Error != nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "ServerError", "data": nil})
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Found", "data": tools})
	} else {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "RecordNotFound", "data": nil})
	}
}
func GetPins(context *gin.Context) {
	// 查询文章表ispinned
	var articles []model.Articles
	id := context.DefaultQuery("id", "1")
	result := data.DB.Where(&model.Articles{ValidFlag: true, IsPinned: true}).Where("club_id = ?", id).Find(&articles)
	if result.Error != nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "ServerError", "data": nil})
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Found", "data": articles})
	} else {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "RecordNotFound", "data": nil})
	}
}
func GetArticles(context *gin.Context) {
	var articles []model.Articles
	id := context.DefaultQuery("id", "1")
	count := context.DefaultQuery("count", "0")
	result := data.DB.Where(&model.Articles{ValidFlag: true}).Where("club_id = ?", id).Find(&articles)
	// todo:
	// 查询文章表，再关联用户表获取作者id名字头像src
	// 根据count裁剪一下
	fmt.Println("clubs: ", articles, count)
	if result.Error != nil {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "ServerError", "data": nil})
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Found", "data": articles})
	} else {
		context.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "RecordNotFound", "data": nil})
	}
}
