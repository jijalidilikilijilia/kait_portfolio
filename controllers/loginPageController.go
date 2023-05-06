package controllers

import (
	"errors"
	"kait_portfolio/database"
	"kait_portfolio/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPageGetController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func LoginPagePostController(ctx *gin.Context) {
	log := ctx.PostForm("login")
	pwd := ctx.PostForm("password")

	var student models.Student

	if err := ctx.ShouldBind(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := authenticate(log, pwd, &student); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	ctx.HTML(http.StatusOK, "home.html", nil)
}

func authenticate(login string, password string, dbStudent *models.Student) error {
	db := database.DB

	if err := db.Table("kait_portfolio.student").Where("login = ?", login).First(&dbStudent).Error; err != nil {
		return err
	}

	if dbStudent.Password != password {
		return errors.New("invalid password")
	}

	return nil
}
