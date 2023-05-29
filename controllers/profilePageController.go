package controllers

import (
	"kait_portfolio/database"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userInfoToShow struct {
	Full_name       string
	Age             int
	Group_name      string
	Cumpus_name     string
	Curator_name    string
	Speciality_name string
	Description     string
	User_photo      []byte
}

func ProfilePageGetController(ctx *gin.Context) {
	userInfoToShow := getUserData(ctx)

	ctx.HTML(http.StatusOK, "profile.html", gin.H{
		"user": userInfoToShow,
	})
}

func getUserData(ctx *gin.Context) userInfoToShow {
	session := sessions.Default(ctx)
	db := database.DB
	var userInfo userInfoToShow

	user_id := session.Get("user_id").(uint)

	err := db.Table("kait_portfolio.student").
		Select("student.full_name", "age", "groups.group_name", "cumpus.cumpus_name", "curators.full_name as curator_name", "specialities.speciality_name", "description", "student.user_photo").
		Joins("JOIN kait_portfolio.groups ON student.group_id = groups.id").
		Joins("JOIN kait_portfolio.cumpus ON student.cumpus_id = cumpus.id").
		Joins("JOIN kait_portfolio.curators ON groups.curator_id = curators.id").
		Joins("JOIN kait_portfolio.specialities ON student.speciality_id = specialities.id").
		Where("student.id = ?", user_id).
		Scan(&userInfo).Error

	if err != nil {
		ctx.Abort()
	}

	userInfo = userInfoToShow{
		Full_name:       userInfo.Full_name,
		Age:             userInfo.Age,
		Group_name:      userInfo.Group_name,
		Cumpus_name:     userInfo.Cumpus_name,
		Curator_name:    userInfo.Curator_name,
		Speciality_name: userInfo.Speciality_name,
		Description:     userInfo.Description,
		User_photo:      userInfo.User_photo,
	}

	return userInfo
}

func ProfilePagePostController(ctx *gin.Context) {
	// Возможно в дальнейшем будут проблемы. Сейчас это ответ на запрос от /edit_profile
	ctx.Redirect(http.StatusFound, "/profile")
}
