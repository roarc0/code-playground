package graph

import (
	"context"
	"time"

	"github.com/roarc0/jobs/database"
)

type Resolver struct {
	db              *database.DB
	timeoutDuration time.Duration
}

func NewResolver(connectionURI string) (*Resolver, error) {
	var db, err = database.Connect(context.Background(), connectionURI)
	if err != nil {
		return nil, err
	}
	return &Resolver{
		db:              db,
		timeoutDuration: 30 * time.Second,
	}, nil
}
