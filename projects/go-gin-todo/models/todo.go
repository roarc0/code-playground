package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title" binding:"required,customTitleCheck"`
	Description string `json:"description" binding:"required,customDescriptionCheck"`
}

func CustomTitleCheck(fl validator.FieldLevel) bool {
	title := fl.Field().String()
	return len(title) >= 5 // Example: Minimum title length
}

func CustomDescriptionCheck(fl validator.FieldLevel) bool {
	description := fl.Field().String()
	return len(description) >= 10 // Example: Minimum description length
}
