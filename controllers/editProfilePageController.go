package controllers

import (
	"io"
	"kait_portfolio/database"
	"kait_portfolio/database/models"
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
	updateStudentDesc(ctx)
	updateStudentPhoto(ctx)

	ctx.Redirect(http.StatusFound, "/profile")
}

func updateStudentDesc(ctx *gin.Context) {
	session := sessions.Default(ctx)
	newDesc := ctx.PostForm("description")
	db := database.DB
	user_id := session.Get("user_id").(uint)

	result := db.Table("kait_portfolio.student").Where("id = ?", user_id).Update("description", newDesc)
	if result.Error != nil {
		panic("ОШИБКА В ЗАПРОСЕ АПДЕЙТ ДЕСКРИПТИОН))")
	}
}

func updateStudentPhoto(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user_id := session.Get("user_id").(uint)
	db := database.DB
	var student models.Student

	file, err := ctx.FormFile("profile-picture")

	if err != nil {
		log.Println("error: ", err)
	}

	if file == nil {
		return
	}

	src, err := file.Open()
	if err != nil {
		log.Println("error: ", err)
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		log.Println("error", err)
	}

	if err := db.Table("kait_portfolio.student").First(&student, user_id).Error; err != nil {
		log.Println(err)
	}

	student.Photo = data

	if err := db.Table("kait_portfolio.student").Save(&student).Error; err != nil {
		log.Println(err)
	}
}
