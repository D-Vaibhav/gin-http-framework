package main

import (
	"vaibhav/try-try-gg/controller"
	"vaibhav/try-try-gg/middlewares"
	"vaibhav/try-try-gg/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	// changing from default to new so that we can implement our own Logger() and Recovery()
	server := gin.New()

	// MIDDLEWARES ie. default in Default()
	server.Use(gin.Recovery(), middlewares.Logger()) // self implemented Logger()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":7000")
}
