package models

import "time"

type File struct {
	ID         string    `bson:"__id,omitempty" json:"id"`
	Filename   string    `bson:"filename" json:"filename"`
	Path       string    `bson:"path" json:"path"`
	Size       int64     `bson:"size" json:"size"`
	MimeType   string	 `bson:"minType" json:"mimeType"`
	UploadedAt time.Time `bson:"uploadedAt" json:"uploadedAt"`
}
