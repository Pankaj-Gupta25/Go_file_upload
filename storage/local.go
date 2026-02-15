package storage

import (
	"fmt"
	"mime/multipart"
	"os"
	"time"
)

func SaveUploadedFile(file *multipart.FileHeader) (string, error) {
	os.MkdirAll("uploads", os.ModePerm)
	unique := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	path := "uploads/" + unique
	return path, nil
}
