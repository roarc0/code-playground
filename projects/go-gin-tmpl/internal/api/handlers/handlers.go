package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/roarc0/go-gin-tmpl/internal/database/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handlers contains all the handlers of the application
type Handlers struct {
	BlogHandler *gin.Engine
}

func NewHandlers(db *mongo.Database) *Handlers {
	ar := repositories.NewArticleRepository(db, "articles")

	return &Handlers{
		BlogHandler: blogHandler(ar),
	}
}
