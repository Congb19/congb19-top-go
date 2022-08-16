package main

import (
	"congb19-top-go/src/controller"
	"congb19-top-go/src/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

const (
	version = "v1.0.0"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
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
	router.Use(Cors())

	api := router.Group("/api/public")
	{

		api.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello GO !",
			})
		})
		api.GET("/getHappinessList", controller.GetHappinessList)
		// api.POST("/postInfo", controller.postInfo)

		api.GET("/getClubs", controller.GetClubs)
		api.GET("/getTools", controller.GetTools)
		api.GET("/getPins", controller.GetPins)
		api.GET("/getArticles", controller.GetArticles)
	}

	router.Run()
}
