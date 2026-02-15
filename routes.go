package main

import "github.com/gin-gonic/gin"

func setupRoutes() *gin.Engine {
	r := gin.New()
	r.GET("/", home)
	r.POST("/file", handler.Upload)
}
