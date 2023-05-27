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
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/hello", controllers.Index)
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
