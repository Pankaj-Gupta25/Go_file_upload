package main

import (
	"github.com/Pankaj-Gupta25/Go_file_upload/handler"
	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	r := gin.New()
	r.GET("/", handler.Home)
	r.POST("/upload", handler.Upload)
	r.Static("/file","./uploads")

	return r
}
