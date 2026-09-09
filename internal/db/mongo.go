package db

import (
	"context"
	"fmt"
	"go-auth/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func Connect(ctx context.Context, cfg config.Config) (*Mongo, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Mongo connection failed: %w", err)
	}

	// Ping the primary to verify the connection
	pingCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if err = client.Ping(pingCtx, nil); err != nil {
		return nil, fmt.Errorf("Mongo ping failed: %w", err)
	}

	database := client.Database(cfg.MongoDatabase)

	return &Mongo{
		Client:   client,
		Database: database,
	}, nil
}
