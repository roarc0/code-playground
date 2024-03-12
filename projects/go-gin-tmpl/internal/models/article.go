package models

import (
	"time"
)

// Article represents the content of a blog article
type Article struct {
	ID      string
	Title   string `form:"title" binding:"required,gte=1,lte=100"`
	Content string `form:"content" binding:"required,gte=10"`
	Date    time.Time
}
