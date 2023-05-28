package routes

import (
	"kait_portfolio/controllers"
	"kait_portfolio/database"
	"kait_portfolio/middleware"

	"github.com/gin-gonic/gin"
)

func AboutRoute(rg *gin.RouterGroup) {
	rg.Use()
	{
		rg.GET("/about", middleware.AuthMiddleware(database.DB), controllers.AboutPageController)
	}
}
