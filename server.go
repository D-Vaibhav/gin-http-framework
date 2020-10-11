package main

import (
	"io"
	"os"
	"vaibhav/try-try-gg/controller"
	"vaibhav/try-try-gg/middlewares"
	"vaibhav/try-try-gg/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

// create a new gin.log file which records all user requests
func setupLogOutput() {
	// outputLogFile, logError
	f, _ := os.Create("gin.log") // file wil get created
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	// changing from default to new so that we can implement our own Logger() and Recovery()
	server := gin.New()

	// MIDDLEWARES ie. default in Default()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":7000")
}
