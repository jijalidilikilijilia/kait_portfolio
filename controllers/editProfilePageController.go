package controllers

import (
	"kait_portfolio/database"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func EditProfileGetController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "edit_profile.html", nil)
}

func EditProfilePostController(ctx *gin.Context) {
	session := sessions.Default(ctx)
	newDesc := ctx.PostForm("description")
	db := database.DB
	user_id := session.Get("user_id").(uint)

	result := db.Table("kait_portfolio.student").Where("id = ?", user_id).Update("description", newDesc)
	if result.Error != nil {
		panic("ОШИБКА В ЗАПРОСЕ АПДЕЙТ ДЕСКРИПТИОН))")
	}

	ctx.Redirect(http.StatusFound, "/profile")
}
