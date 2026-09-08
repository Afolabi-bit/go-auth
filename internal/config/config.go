package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI      string
	MongoDatabase string
	JWTSecret     string
}

func LoadConfig() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		MongoURI:      strings.TrimSpace(os.Getenv("MONGO_URI")),
		MongoDatabase: strings.TrimSpace(os.Getenv("MONGO_DB_NAME")),
		JWTSecret:     strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}

	if cfg.MongoURI == "" {
		return Config{}, fmt.Errorf("Environment Variable MONGO_URI not set")
	}
	if cfg.MongoDatabase == "" {
		return Config{}, fmt.Errorf("Environment Variable MONGO_DB_NAME not set")
	}
	if cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("Environment Variable JWT_SECRET not set")
	}

	return cfg, nil
}
