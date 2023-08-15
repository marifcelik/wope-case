package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
