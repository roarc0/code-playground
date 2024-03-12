package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/roarc0/go-gin-tmpl/internal/models"
)

type ArticleRepository struct {
	collection *mongo.Collection
}

type dbArticle struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
	Date    time.Time          `bson:"date"`
}

func (b *dbArticle) ToModel() models.Article {
	return models.Article{
		ID:      b.ID.Hex(),
		Title:   b.Title,
		Content: b.Content,
		Date:    b.Date,
	}
}

func NewArticleRepository(db *mongo.Database, collectionName string) *ArticleRepository {
	return &ArticleRepository{collection: db.Collection(collectionName)}
}

// Create adds a new article to the db. If the time is zero it will get the current
func (ar *ArticleRepository) Create(ctx context.Context, article *models.Article) error {
	if article.Date.IsZero() {
		article.Date = time.Now()
	}
	_, err := ar.collection.InsertOne(ctx, article)
	return err
}

func (ar *ArticleRepository) ReadByID(ctx context.Context, id string) (*models.Article, error) {
	pID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var da dbArticle
	err = ar.collection.FindOne(ctx, primitive.M{"_id": pID}).Decode(&da)
	if err != nil {
		return nil, err
	}
	article := da.ToModel()
	return &article, nil
}

func (ar *ArticleRepository) ReadPaged(ctx context.Context, pageSize, pageNumber int) ([]models.Article, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64((pageNumber - 1) * pageSize))
	findOptions.SetSort(primitive.D{{Key: "date", Value: -1}}) // Sort by date descending

	cursor, err := ar.collection.Find(ctx, primitive.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	var articlesDB []*dbArticle
	err = cursor.All(ctx, &articlesDB)
	if err != nil {
		return nil, err
	}

	var articles []models.Article
	for _, p := range articlesDB {
		articles = append(articles, p.ToModel())
	}
	return articles, nil
}

func (ar *ArticleRepository) Update(ctx context.Context, article *models.Article) error {
	pID, err := primitive.ObjectIDFromHex(article.ID)
	if err != nil {
		return err
	}

	_, err = ar.collection.ReplaceOne(ctx, primitive.M{"_id": pID}, article)
	return err
}

func (ar *ArticleRepository) Delete(ctx context.Context, id string) error {
	pID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = ar.collection.DeleteOne(ctx, primitive.M{"_id": pID})
	return err
}
