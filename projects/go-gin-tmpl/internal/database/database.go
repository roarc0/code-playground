package database

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/roarc0/go-gin-tmpl/internal/config"
)

func Connect(cfg *config.Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Debug().Msg("Connecting to MongoDB")

	clientOptions := options.Client().
		ApplyURI(cfg.DBURI()).
		SetMaxConnIdleTime(time.Duration(cfg.DB.MaxIdleConns)).
		SetMaxConnecting(uint64(cfg.DB.MaxOpenConns)).
		SetMaxPoolSize(uint64(cfg.DB.MaxOpenConns))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database("blog")

	return db, nil
}
