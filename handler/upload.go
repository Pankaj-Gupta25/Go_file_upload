package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	// db "command-line-argumentsC:\\Users\\PANKAJ\\OneDrive\\Desktop\\coding\\GO\\Gin\\Day-2\\db\\mongo.go"
	"github.com/Pankaj-Gupta25/Go_file_upload/models"
	// "github.com/Pankaj-Gupta25/Go_file_upload/db"
	db "github.com/Pankaj-Gupta25/Go_file_upload/db"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		// c.String(http.StatusBadRequest,"get form err: %s",err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file required",
		})
		return
	}

	path := "uploads/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		log.Fatal(err)
		return
	}

	meta := new(models.File)

	meta.Filename = file.Filename
	meta.Path = path
	meta.Size = file.Size
	meta.MimeType = file.Header.Get("Content-Type")
	meta.UploadedAt = time.Now()

	db.FileCollection.InsertOne(context.TODO(),meta)
	
	c.JSON(http.StatusOK,gin.H{
		"message":"uploaded",
		"url":"/files/"+file.Filename,
	})

}
