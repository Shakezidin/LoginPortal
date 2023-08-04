package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/shaikh_zidhin/controllers"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	r.Use(controllers.NoCache())

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/", controllers.LandingPage)
	r.GET("/login", controllers.ShowLoginPage)
	r.POST("/login/user", controllers.UserLogin)
	r.GET("/home", controllers.ShowHomePage)
	r.GET("/logout", controllers.Logout)

	r.Run("localhost:8080")
}
