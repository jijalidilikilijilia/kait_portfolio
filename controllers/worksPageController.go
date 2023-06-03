package controllers

import (
	"kait_portfolio/database"
	"kait_portfolio/database/models"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func WorksPageGetController(ctx *gin.Context) {
	works := []models.StudentWorks{}
	db := database.DB
	session := sessions.Default(ctx)
	user_id := session.Get("user_id").(uint)

	err := db.Table(`kait_portfolio."studentWorks"`).Where("student_id = ?", user_id).Find(&works).Error

	if err != nil {
		log.Println("error select works: ", err)
	}

	if len(works) == 0 {
		ctx.HTML(http.StatusOK, "works.html", gin.H{
			"message": "Вы ещё ничего не загрузили",
		})
	} else {
		ctx.HTML(http.StatusOK, "works.html", gin.H{
			"works": works,
		})
	}
}

func WorksPagePostController(ctx *gin.Context) {
	// нужно будет сделать редирект чтобы обновить список работ
}
