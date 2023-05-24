package routes

import (
	"kait_portfolio/controllers"

	"github.com/gin-gonic/gin"
)

func LogoutRoute(rg *gin.RouterGroup) {
	rg.Use()
	{
		rg.GET("/logout", controllers.LogoutController)
	}
}
