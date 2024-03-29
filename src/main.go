package main

import (
	"congb19-top-go/src/controller"
	"congb19-top-go/src/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/unrolled/secure"
	"log"
	"net/http"
)

//const (
//	version = "v1.0.0"
//)

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
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, session, Content-Type")
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

func loadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost: 8003",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			fmt.Println(err)
			return
		}
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
	test := router.Group("")
	{
		test.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello GO !",
			})
		})
	}
	app := router.Group("/api/app")
	{
		app.GET("/getHappinessList", controller.GetHappinessList)
		// api.POST("/postInfo", controller.postInfo)

		app.GET("/getClubs", controller.GetClubs)
		app.GET("/getTools", controller.GetTools)
		app.GET("/getPins", controller.GetPins)
		app.GET("/getArticles", controller.GetArticles)
	}
	kbn := router.Group("/api/kbn")
	{
		kbn.GET("/getHappinessList", controller.GetHappinessList)
		kbn.POST("/postKbn", controller.PostKbn)

	}

	// 开发环境下使用http
	err = router.Run(":8003")

	// 部署时使用https
	//router.Use(loadTls())
	//err = router.RunTLS(":8003", "./fullchain.pem", "./privkey.key")
	//if err != nil {
	//	return
	//}
}
