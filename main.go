package main

import (
	"go-auth-backend/controllers"
	"go-auth-backend/database"
	"go-auth-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// init database
	database.Connect("root:password@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()

	// init router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
