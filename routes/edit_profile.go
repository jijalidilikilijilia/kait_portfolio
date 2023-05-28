package routes

import (
	"kait_portfolio/controllers"
	"kait_portfolio/database"
	"kait_portfolio/middleware"

	"github.com/gin-gonic/gin"
)

func EditProfileRoute(rg *gin.RouterGroup) {
	rg.Use()
	{
		rg.GET("/edit_profile", middleware.AuthMiddleware(database.DB), controllers.EditProfileGetController)
		rg.POST("/edit_profile", middleware.AuthMiddleware(database.DB), controllers.EditProfilePostController)
	}
}
