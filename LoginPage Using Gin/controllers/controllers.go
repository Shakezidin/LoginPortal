package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shaikh_zidhin/models"
)

func NoCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache,no-store, must-revalidate")
		c.Header("Expires", "0")
		c.Next()
	}
}

func ShowLoginPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	if username != nil {
		c.HTML(http.StatusOK, "homepage.html", nil)
		return
	}
	c.HTML(http.StatusOK, "login.html", nil)
}

func LandingPage(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/login")
}

var data models.Data

func UserLogin(c *gin.Context) {
	data.Username = c.PostForm("username")
	data.Password = c.PostForm("password")

	if Validation(data.Username, data.Password) {
		session := sessions.Default(c)
		session.Set("username", data.Username)
		session.Save()
		c.Redirect(http.StatusSeeOther, "/home")
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"ErrorMessage": "Invalid username or password.",
	})
}

func Validation(username, password string) bool {
	return username == "shaikh_zidhin" && password == "1090"
}

func ShowHomePage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	if username == nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	c.HTML(http.StatusOK, "homepage.html", gin.H{
		"Username": username,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("username")
	session.Save()

	c.Redirect(http.StatusSeeOther, "/login")
}
