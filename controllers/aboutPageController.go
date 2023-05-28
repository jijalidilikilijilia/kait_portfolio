package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutPageController(ctx *gin.Context) {
	ctx.HTML(http.StatusFound, "error_page.html", nil)
}
