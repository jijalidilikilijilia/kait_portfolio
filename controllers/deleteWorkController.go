package controllers

import (
	"fmt"
	"kait_portfolio/database"
	"kait_portfolio/database/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DeleteWorkGetController(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/profile")
}

func DeleteWorkByIdGetController(ctx *gin.Context) {
	studentWorkID := ctx.Param("work_id")
	session := sessions.Default(ctx)
	user_id := session.Get("user_id").(uint)
	db := database.DB
	var studentWork models.StudentWorks

	if err := db.Table("kait_portfolio.studentWorks").Delete(&studentWork, studentWorkID).Where("student_id = ?", user_id).Error; err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка удаления записи из базы данных: %s", err.Error()))
		return
	}

	ctx.Redirect(http.StatusFound, "/works")
}

func DeleteWorkByIdPostController(ctx *gin.Context) {

}
