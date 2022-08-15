package main

import (
	"congb19-top-go/src/controller"
	"congb19-top-go/src/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	version = "v1.0.0"
)

func main() {
	gin.SetMode(gin.DebugMode)

	// get config
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")   // 配置文件的路径
	err := viper.ReadInConfig() //找到并读取配置文件
	if err != nil {             // 捕获读取中遇到的error
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// DB
	handler.InitDB()

	// router
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello GO !",
			})
		})
		// api.POST("/postInfo", controller.postInfo)
		api.GET("/getHappinessList", controller.GetHappinessList)
		api.GET("/getClubs", controller.GetClubs)
	}

	router.Run()
}
