package main

import (
	"qrcode-login/controllers"
	"qrcode-login/database"
	"qrcode-login/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("root:root@tcp(db:3306)/qrcode?parseTime=true")
	database.Migrate()

	router := initRouter()

	router.Use(CORS())
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/hello", controllers.Index)
		api.POST("/auth/login", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)

		secured := api.Group("/").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			secured.GET("/auth/user", controllers.UserInfo)
		}
	}

	return router
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
