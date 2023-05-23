package routes

import (
	"kait_portfolio/controllers"
	"kait_portfolio/database"
	"kait_portfolio/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileRoute(rg *gin.RouterGroup) {
	rg.Use()
	{
		rg.GET("/profile", middleware.AuthMiddleware(database.DB), controllers.ProfilePageGetController)
	}
}
