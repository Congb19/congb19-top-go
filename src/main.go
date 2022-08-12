package main

import (
	"congb19-top-go/src/controller"
	"congb19-top-go/src/handler"
	"github.com/gin-gonic/gin"
)

const (
	version = "v1.0.0"
)

func main() {
	gin.SetMode(gin.DebugMode)

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
