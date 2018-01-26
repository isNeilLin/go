package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	Controller "./controllers"
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()

	r.GET("/ping", Controller.Ping)
	r.GET("/user/:name/*action", Controller.User)
	r.GET("/welcome", Controller.Welcome)
	r.POST("/form_post", Controller.FormPost)
	r.POST("/upload", Controller.Upload)

	r.Run()
}
