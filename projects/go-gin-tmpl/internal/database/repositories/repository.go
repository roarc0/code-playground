package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	ArticleRepository *ArticleRepository
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		ArticleRepository: NewArticleRepository(db, "articles"),
	}
}
