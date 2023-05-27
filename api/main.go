package main

import (
	"github.com/gin-gonic/gin"
	"qrcode-login/controllers"
	"qrcode-login/database"
	"qrcode-login/middlewares"
)

func main() {
	database.Connect("root:root@tcp(localhost:3306)/qrcode?parseTime=true")
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
