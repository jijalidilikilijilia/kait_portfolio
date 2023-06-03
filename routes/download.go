package routes

import (
	"kait_portfolio/controllers"
	"kait_portfolio/database"
	"kait_portfolio/middleware"

	"github.com/gin-gonic/gin"
)

func Download(rg *gin.RouterGroup) {
	rg.Use()
	{
		rg.GET("/download/", middleware.AuthMiddleware(database.DB), controllers.DownloadFileGetController)
		rg.GET("/download/:work_id", middleware.AuthMiddleware(database.DB), controllers.DownloadFileByIdController)
		rg.POST("/download/", middleware.AuthMiddleware(database.DB), controllers.DownloadFileGetController)
	}
}
