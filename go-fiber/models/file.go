package models

import "time"

type File struct {
	Dir           string     `json:"dir" gorm:"index"`
	FileExtension string     `json:"file_extension"`
	FileID        string     `json:"file_id"`
	Name          string     `json:"name" gorm:"index"`
	Type          string     `json:"type"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Category      string     `json:"category"`
	ContentType   string     `json:"content_type"`
	Size          int64      `json:"size"`
	Password      string     `json:"password"`
	Url           string     `json:"url" gorm:"_"`
}
