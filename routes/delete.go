package routes

import (
	"kait_portfolio/controllers"
	"kait_portfolio/database"
	"kait_portfolio/middleware"

	"github.com/gin-gonic/gin"
)

func Delete(rg *gin.RouterGroup) {
	rg.Use()
	{
		rg.GET("/delete/", middleware.AuthMiddleware(database.DB), controllers.DeleteWorkGetController)
		rg.GET("/delete/:work_id", middleware.AuthMiddleware(database.DB), controllers.DeleteWorkByIdGetController)
		rg.POST("/delete/", middleware.AuthMiddleware(database.DB), controllers.DeleteWorkByIdPostController)
	}
}
