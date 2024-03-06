package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.comroarc0/go-gin-todo/middlewares"
	"github.comroarc0/go-gin-todo/models"
)

func NewTodoController(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.LoggerMiddleware())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("customTitleCheck", models.CustomTitleCheck)
		v.RegisterValidation("customDescriptionCheck", models.CustomDescriptionCheck)
	}

	router.POST("/todos", func(c *gin.Context) {
		var todo models.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			return
		}
		result := db.Create(&todo)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Db error"})
			return
		}
		c.JSON(http.StatusOK, todo)
	})

	router.GET("/todos", func(c *gin.Context) {
		var todos []models.Todo
		db.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	router.GET("/todos/:id", func(c *gin.Context) {
		var todo models.Todo
		todoID := c.Param("id")
		result := db.First(&todo, todoID)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusOK, todo)
	})

	router.PUT("/todos/:id", func(c *gin.Context) {
		var todo models.Todo
		todoID := c.Param("id")
		result := db.First(&todo, todoID)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		var updatedTodo models.Todo
		if err := c.ShouldBindJSON(&updatedTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			return
		}

		todo.Title = updatedTodo.Title
		todo.Description = updatedTodo.Description
		result = db.Save(&todo)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Db error"})
			return
		}
		c.JSON(http.StatusOK, todo)
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		var todo models.Todo
		todoID := c.Param("id")
		result := db.First(&todo, todoID)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		result = db.Delete(&todo)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Db error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Todo with ID %s deleted", todoID)})
	})

	return router
}
