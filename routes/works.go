package routes

import (
	"kait_portfolio/controllers"
	"kait_portfolio/database"
	"kait_portfolio/middleware"

	"github.com/gin-gonic/gin"
)

func WorksRoute(rg *gin.RouterGroup) {
	rg.Use()
	{
		rg.GET("/works", middleware.AuthMiddleware(database.DB), controllers.WorksPageGetController)
		rg.POST("/works", middleware.AuthMiddleware(database.DB), controllers.WorksPagePostController)
	}
}
