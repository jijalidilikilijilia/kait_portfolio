package main

import (
	"kait_portfolio/database"
	"kait_portfolio/initializers"
	"kait_portfolio/routes"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	rootPath := router.Group("/")

	router.LoadHTMLGlob("templates/*")
	router.Static("/public/css", "./public/css")
	router.Static("/public/img", "./public/img")
	router.Static("/public/js", "./public/js")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	routes.LoginRoute(rootPath)
	routes.ProfileRoute(rootPath)
	routes.LogoutRoute(rootPath)
	routes.AboutRoute(rootPath)
	routes.EditProfileRoute(rootPath)
	routes.WorksRoute(rootPath)

	router.Run(":8000")
}
