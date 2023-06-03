package controllers

import (
	"fmt"
	"kait_portfolio/database"
	"kait_portfolio/database/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DownloadFileGetController(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/profile")
}

func DownloadFileByIdController(ctx *gin.Context) {
	studentWorkID := ctx.Param("work_id")
	session := sessions.Default(ctx)
	user_id := session.Get("user_id").(uint)
	db := database.DB
	var studentWork models.StudentWorks

	// Получение записи StudentWork из базы данных по идентификатору
	if err := db.Table("kait_portfolio.studentWorks").First(&studentWork, studentWorkID).Where("student_id = ?", user_id).Error; err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка получения записи из базы данных: %s", err.Error()))
		return
	}

	// Установка заголовков HTTP для скачивания файла
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", studentWork.File_name))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Length", fmt.Sprintf("%d", len(studentWork.File)))

	// Отправка содержимого файла в ответе HTTP
	ctx.Writer.Write(studentWork.File)
}

func DownloadFilePostController(ctx *gin.Context) {
	work_id := ctx.Param("work_id")
	ctx.String(200, "ID = ", work_id)
}
