package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roarc0/go-gin-tmpl/assets"
	"github.com/roarc0/go-gin-tmpl/database"
	"github.com/roarc0/go-gin-tmpl/models"
	"github.com/rs/zerolog/log"
)

// Blog represents a whole blog page
type Blog struct {
	Title    string
	Articles []models.Article
	Error    error
}

func getBlogRouter(articleService *database.ArticleService) *gin.Engine {
	router := gin.Default()
	router.SetHTMLTemplate(assets.Templates())

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("len", bookableDate)
	// }

	staticFs, err := assets.Static()
	if err != nil {
		log.Fatal().Err(err).Msg("could not retrieve static files")
	}
	router.StaticFS("/public", http.FS(staticFs))

	router.GET("/", func(c *gin.Context) {
		var postError error = nil
		if msg, ok := c.GetQuery("error"); ok {
			postError = errors.New(msg)
		}

		articles, err := articleService.ReadPaged(c.Request.Context(), 10, 1)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.HTML(http.StatusOK, "index.tmpl", Blog{
			Title:    "My Blog",
			Articles: articles,
			Error:    postError,
		})
	})

	router.POST("/create_article", func(c *gin.Context) {
		var article models.Article
		if err := c.ShouldBind(&article); err != nil {
			c.Redirect(http.StatusSeeOther, "/?error=Error creating post")
			return
		}

		if err := articleService.Create(c.Request.Context(), &article); err != nil {
			c.Redirect(http.StatusSeeOther, "/?error=Error creating post")
			return
		}

		c.Redirect(http.StatusSeeOther, "/")
	})

	return router
}
