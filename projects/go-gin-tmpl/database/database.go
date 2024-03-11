package database

import (
	"context"
	"time"

	"github.com/roarc0/go-gin-tmpl/models"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArticleService struct {
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

func NewArticleService(ctx context.Context, connectionString, databaseName, collectionName string) (*ArticleService, error) {
	log.Debug().Msg("Connecting to MongoDB")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	log.Debug().Msg("Pinging MongoDB")
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database(databaseName).Collection(collectionName)

	return &ArticleService{collection: collection}, nil
}

// Create adds a new article to the db. If the time is zero it will get the current
func (s *ArticleService) Create(ctx context.Context, article *models.Article) error {
	if article.Date.IsZero() {
		article.Date = time.Now()
	}
	_, err := s.collection.InsertOne(ctx, article)
	return err
}

func (s *ArticleService) ReadByID(ctx context.Context, id string) (*models.Article, error) {
	pID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var da dbArticle
	err = s.collection.FindOne(ctx, primitive.M{"_id": pID}).Decode(&da)
	if err != nil {
		return nil, err
	}
	article := da.ToModel()
	return &article, nil
}

func (s *ArticleService) ReadPaged(ctx context.Context, pageSize, pageNumber int) ([]models.Article, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64((pageNumber - 1) * pageSize))
	findOptions.SetSort(primitive.D{{Key: "date", Value: -1}}) // Sort by date descending

	cursor, err := s.collection.Find(ctx, primitive.M{}, findOptions)
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

func (s *ArticleService) Update(ctx context.Context, article *models.Article) error {
	pID, err := primitive.ObjectIDFromHex(article.ID)
	if err != nil {
		return err
	}

	_, err = s.collection.ReplaceOne(ctx, primitive.M{"_id": pID}, article)
	return err
}

func (s *ArticleService) Delete(ctx context.Context, id string) error {
	pID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(ctx, primitive.M{"_id": pID})
	return err
}
