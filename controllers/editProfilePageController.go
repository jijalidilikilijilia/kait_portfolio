package controllers

import (
	"kait_portfolio/database"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func EditProfileGetController(ctx *gin.Context) {
	db := database.DB
	session := sessions.Default(ctx)
	user_id := session.Get("user_id").(uint)
	var description string

	err := db.Table("kait_portfolio.student").Select("description").Where("id = ?", user_id).Scan(&description).Error
	if err != nil {
		log.Println("АШИБКА ЗАПРОСА ЗАМЕНА ДЕСК")
	}

	ctx.HTML(http.StatusOK, "edit_profile.html", gin.H{
		"description": description,
	})
}

func EditProfilePostController(ctx *gin.Context) {
	session := sessions.Default(ctx)
	newDesc := ctx.PostForm("description")
	user_photo := getStudentPhotoData(ctx)
	db := database.DB
	user_id := session.Get("user_id").(uint)

	result := db.Table("kait_portfolio.student").Where("id = ?", user_id).Update("description", newDesc).Update("user_photo", user_photo)
	if result.Error != nil {
		panic("ОШИБКА В ЗАПРОСЕ АПДЕЙТ ДЕСКРИПТИОН))")
	}

	ctx.Redirect(http.StatusFound, "/profile")
}

func getStudentPhotoData(ctx *gin.Context) []byte {
	file, err := ctx.FormFile("profile-picture")
	if err != nil {
		ctx.String(400, "error file downloadj")
		ctx.Abort()
	}

	src, err := file.Open()
	if err != nil {
		ctx.String(400, "error file open")
		ctx.Abort()
	}
	defer src.Close()

	data := make([]byte, file.Size)
	if _, err := src.Read(data); err != nil {
		ctx.String(400, "error reading file data")
		ctx.Abort()
	}

	return data
}
