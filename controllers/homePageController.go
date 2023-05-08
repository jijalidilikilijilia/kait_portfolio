package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePageGetController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Домашнаяя страница",
	})
}

func HomePagePostController(ctx *gin.Context) {

}
