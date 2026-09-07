package app

import (
	"context"
	"fmt"
	"go-auth/internal/config"
	"go-auth/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Config        config.Config
	MongoClient   *mongo.Client
	MongoDatabase *mongo.Database
}

func NewApp(ctx context.Context) (*App, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("Failed to load config: %w", err)
	}

	mongoCli, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to MongoDB: %w", err)
	}

	return &App{
		Config:        cfg,
		MongoClient:   mongoCli.Client,
		MongoDatabase: mongoCli.Database,
	}, nil
}

func (a *App) Close(ctx context.Context) error {
	if a.MongoClient == nil {
		return nil
	}
	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("Failed to disconnect from MongoDB: %w", err)
	}
	return nil
}
