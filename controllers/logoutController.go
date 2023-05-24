package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogoutController(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Set("authenticated", false)

	session.Save()

	ctx.Redirect(http.StatusFound, "/login")
}
