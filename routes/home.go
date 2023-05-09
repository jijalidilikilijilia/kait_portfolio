package routes

import (
	"kait_portfolio/controllers"
	"kait_portfolio/database"
	"kait_portfolio/middleware"

	"github.com/gin-gonic/gin"
)

func HomeRoute(rg *gin.RouterGroup) {
	rg.Use()
	{
		rg.GET("/home", middleware.AuthMiddleware(database.DB), controllers.HomePageGetController)
		rg.POST("/home", controllers.HomePagePostController)
	}
}
