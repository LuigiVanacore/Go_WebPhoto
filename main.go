package main

import (
	"Go_WebPhoto/controllers"
	"Go_WebPhoto/views"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	homeView    *views.View
	contactView *views.View
	signupView  *views.View
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home", nil)
}

func contact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact", nil)
}

func signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup", nil)
}

func main() {
	homeView = views.NewView("home", "static/templates/home.gohtml")
	contactView = views.NewView("contact", "static/templates/contact.gohtml")
	usersC := controllers.NewUsers()
	r := multitemplate.NewRenderer()

	r.Add("home", homeView.Template)
	r.Add("contact", contactView.Template)
	r.Add("signup", usersC.NewView.Template)
	router := gin.Default()
	router.HTMLRender = r
	router.GET("/contact", contact)
	router.GET("/signup", usersC.New)
	router.POST("/signup", usersC.Create)
	router.GET("/", home)
	router.Run()
}
