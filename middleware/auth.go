package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		if session.Get("authenticated") != true {
			ctx.Redirect(http.StatusFound, "/login")
			ctx.Abort()
			return
		}

		// if auth go next
		ctx.Next()
	}
}
