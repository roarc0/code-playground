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

const (
	blogTitle = "My Blog"
)

// Page is used to understand which page we want to render
type Page struct {
	PageType  string
	BlogTitle string
	Error     error
}

// IndexPage represents a whole blog page
type IndexPage struct {
	Page
	Articles []models.Article
}

// ArticlePage represents a whole blog page
type ArticlePage struct {
	Page
	models.Article
}

func getBlogRouter(articleService *database.ArticleService) *gin.Engine {
	router := gin.Default()
	router.SetHTMLTemplate(assets.Templates())

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
			return
		}

		c.HTML(http.StatusOK, "index.tmpl", IndexPage{
			Page: Page{
				BlogTitle: blogTitle,
				Error:     postError,
				PageType:  "index",
			},
			Articles: articles,
		})
	})

	router.GET("/article/:id", func(c *gin.Context) {
		type Identifiable struct {
			ID string `uri:"id"`
		}

		var id Identifiable
		err := c.ShouldBindUri(&id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		log.Info().Any("id", id.ID).Msg("id")

		article, err := articleService.ReadByID(c.Request.Context(), id.ID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.HTML(http.StatusOK, "index.tmpl", ArticlePage{
			Page: Page{
				BlogTitle: blogTitle,
				PageType:  "article",
			},
			Article: *article,
		})
	})

	router.POST("/create_article", func(c *gin.Context) {
		var article models.Article
		if err := c.ShouldBind(&article); err != nil {
			c.Redirect(http.StatusSeeOther, "/?error=Error: invalid post")
			return
		}

		if err := articleService.Create(c.Request.Context(), &article); err != nil {
			c.Redirect(http.StatusSeeOther, "/?error=Error: creating post")
			return
		}

		c.Redirect(http.StatusSeeOther, "/")
	})

	return router
}
