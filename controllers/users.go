package controllers

import (
	"Go_WebPhoto/views"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("user", "static/templates/signup.gohtml"),
	}
}

func (u *Users) New(c *gin.Context) {
	c.HTML(http.StatusOK, "signup", nil)
}

func (u *Users) Create(c *gin.Context) {

	fmt.Fprintln(c.Writer, c.PostForm("email"))
	fmt.Fprintln(c.Writer, c.PostForm("password"))
}
