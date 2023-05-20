package controllers

import (
	"errors"
	"kait_portfolio/database"
	"kait_portfolio/database/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginPageGetController(ctx *gin.Context) {
	session := sessions.Default(ctx)

	if session.Get("authenticated") == true {
		ctx.Redirect(http.StatusFound, "/home")
	}

	ctx.HTML(http.StatusFound, "login.html", gin.H{
		"Error": false,
	})
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
		ctx.HTML(http.StatusBadRequest, "login.html", gin.H{
			"Error":        true,
			"ErrorMessage": "Неправильный логин или пароль",
		})
		return
	}

	session := sessions.Default(ctx)
	session.Set("authenticated", true)
	session.Set("username", log)
	session.Save()

	ctx.Redirect(http.StatusFound, "/home")
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
