package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerFunc(c *gin.Context) {
	fmt.Fprint(c.Writer, "<h1>Welcome to my awesome site!</h1>")

	c.Header("Content-Type", "text/html")
	if c.Request.URL.Path == "/" {
		fmt.Fprint(c.Writer, "<h1>Welcome to my awesome site!</h1>")
	} else if c.Request.URL.Path == "/contact" {
		fmt.Fprint(c.Writer, "To get in touch, please send an email "+
			"to <a href=\"mailto:luigi.vanacore@outlook.com\">"+
			"luigi.vanacore@outlook.com</a>.")
	}
}

func hello(c *gin.Context) {
	var name = "Luigi"
	c.HTML(http.StatusOK, "hello.gohtml", gin.H{
		"Name": name,
	})
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.gohtml", nil)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handlerFunc)
	router.GET("/hello", hello)
	router.GET("/home", home)
	router.Run()
}
