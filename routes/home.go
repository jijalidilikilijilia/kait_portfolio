package routes

import (
	"kait_portfolio/controllers"

	"github.com/gin-gonic/gin"
)

func HomeRoute(rg *gin.RouterGroup) {
	rg.Use()
	{
		rg.GET("/home", controllers.HomePageGetController)
		rg.POST("/home", controllers.HomePagePostController)
	}
}
