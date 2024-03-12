package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/roarc0/go-gin-tmpl/internal/assets"
	"github.com/roarc0/go-gin-tmpl/internal/database/repositories"
	"github.com/roarc0/go-gin-tmpl/internal/models"
)

const (
	blogTitle = "My Blog"
)

type page struct {
	PageType  string
	BlogTitle string
	Error     error
}

type indexPage struct {
	page
	Articles []models.Article
}

type articlePage struct {
	page
	models.Article
}

func blogHandler(articleRepository *repositories.ArticleRepository) *gin.Engine {
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

		articles, err := articleRepository.ReadPaged(c.Request.Context(), 10, 1)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.HTML(http.StatusOK, "index.tmpl", indexPage{
			page: page{
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

		article, err := articleRepository.ReadByID(c.Request.Context(), id.ID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.HTML(http.StatusOK, "index.tmpl", articlePage{
			page: page{
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

		if err := articleRepository.Create(c.Request.Context(), &article); err != nil {
			c.Redirect(http.StatusSeeOther, "/?error=Error: creating post")
			return
		}

		c.Redirect(http.StatusSeeOther, "/")
	})

	return router
}
