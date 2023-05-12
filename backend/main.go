package main

import (
	"backend/controllers"
	"backend/database"
	"backend/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Migrate()

	router := InitRouter()
	router.Run(":8080")
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
