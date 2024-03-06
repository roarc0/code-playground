package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.comroarc0/go-gin-todo/models"
)

// OpenDB opens db and migrates it
func OpenDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
