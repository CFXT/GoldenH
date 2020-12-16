package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	app.LoadHTMLGlob("templates/*")

	//define Routes

	//Host and teplate the home page
	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome Back To RAK19",
			"time":  time.Now(),
		})
	})

	//define login route
	app.POST("api/v2/login", Login)

	app.Run(":8080")

}
