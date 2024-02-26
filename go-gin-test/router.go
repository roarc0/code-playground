package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMainRouter() *gin.Engine {
	router := gin.Default()
	router.Use(LoggerMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	router.GET("/bye", func(c *gin.Context) {
		c.String(http.StatusOK, "Goodbye, World!")
	})

	// Route using the UserController
	userController := &UserController{}
	router.GET("/users/:id", userController.GetUserInfo)

	// Route with query parameters
	router.GET("/search", func(c *gin.Context) {
		query := c.DefaultQuery("q", "default-value")
		c.String(200, "Search query: "+query)
	})

	public := router.Group("/public")
	{
		public.GET("/info", func(c *gin.Context) {
			c.String(200, "Public information")
		})
		public.GET("/products", func(c *gin.Context) {
			c.String(200, "Public product list")
		})
	}

	authGroup := router.Group("/api")
	authGroup.Use(AuthMiddleware())
	{
		authGroup.GET("/data", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Authenticated and authorized!"})
		})
	}

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.POST("/send-file", func(c *gin.Context) {
		c.mul
	})

	return router
}
