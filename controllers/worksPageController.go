package controllers

import (
	"fmt"
	"kait_portfolio/database"
	"kait_portfolio/database/models"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func WorksPageGetController(ctx *gin.Context) {
	works := []models.StudentWorks{}
	db := database.DB
	session := sessions.Default(ctx)
	user_id := session.Get("user_id").(uint)

	err := db.Table(`kait_portfolio."studentWorks"`).Where("student_id = ?", user_id).Find(&works).Scan(&works).Error

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
	session := sessions.Default(ctx)
	user_id := session.Get("user_id").(uint)
	db := database.DB
	var studentWork models.StudentWorks

	file, err := ctx.FormFile("work")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("Ошибка загрузки файла: %s", err.Error()))
		return
	}

	// Открытие загруженного файла
	src, err := file.Open()
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка открытия файла: %s", err.Error()))
		return
	}
	defer src.Close()

	// Чтение содержимого файла
	fileBytes, err := getFileBytes(src)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка чтения файла: %s", err.Error()))
		return
	}

	// Создание новой записи StudentWork
	studentWork = models.StudentWorks{
		Student_id:  user_id,
		File_name:   file.Filename,
		Upload_date: time.Now().Format("02-06-2006"),
		File:        fileBytes,
	}

	// Сохранение записи в базе данных
	if err := db.Table("kait_portfolio.studentWorks").Create(&studentWork).Error; err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка сохранения в базе данных: %s", err.Error()))
		return
	}

	log.Println("Файл успешно загружен и сохранен в базе данных")

	ctx.Redirect(http.StatusFound, "/works")
}

func getFileBytes(file multipart.File) ([]byte, error) {
	fileBytes := make([]byte, 50000000) //50мб
	_, err := file.Read(fileBytes)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}
