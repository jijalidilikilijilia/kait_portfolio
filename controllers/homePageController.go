package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userInfoToShow struct {
	Username string
	Age      int
	// 	Group_name   string
	// 	Cumpus_name  string
	// 	Curator_name string
	// 	Speciality   string
	// 	About        string
	Id uint
}

func HomePageGetController(ctx *gin.Context) {
	userInfoToShow := getUserData(ctx)

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"user": userInfoToShow,
	})
}

func HomePagePostController(ctx *gin.Context) {

}

func getUserData(ctx *gin.Context) userInfoToShow {
	session := sessions.Default(ctx)

	username := session.Get("username")
	user_id := session.Get("user_id")
	//user_id := session.Get("user_id")

	if username == nil {
		ctx.HTML(200, "error_page.html", nil)
	}

	userInfo := userInfoToShow{
		Username: username.(string),
		Age:      21,
		Id:       user_id.(uint),
	}

	return userInfo
}
