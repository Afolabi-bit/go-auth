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
	connectCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(connectCtx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Mongo connection failed: %w", err)
	}

	database := client.Database(cfg.MongoDatabase)
	//Ping the primary to verify the connection
	if err = client.Ping(connectCtx, nil); err != nil {
		return nil, fmt.Errorf("Mongo ping failed: %w", err)
	}

	return &Mongo{
		Client:   client,
		Database: database,
	}, nil
}
