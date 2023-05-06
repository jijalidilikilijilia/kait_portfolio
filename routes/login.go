package routes

import (
	"kait_portfolio/controllers"

	"github.com/gin-gonic/gin"
)

func LoginRoute(rg *gin.RouterGroup) {
	rg.Use()
	{
		rg.GET("/login", controllers.LoginPageGetController)
		rg.POST("/login", controllers.LoginPagePostController)
	}
}
