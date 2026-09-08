package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Repo struct {
	column *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		column: db.Collection("users"),
	}
}

func (r *Repo) FindByEmail(ctx context.Context, email string) (User, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	var user User

	filter := bson.M{"email": email}

	err := r.column.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return User{}, mongo.ErrNoDocuments
		}
		return User{}, fmt.Errorf("find user by email failed: %w", err)
	}
	return user, nil
}

func (r *Repo) Create(ctx context.Context, user User) (User, error) {
	res, err := r.column.InsertOne(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("Insert user failed: %w", err)
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("InsertedID user failed and inserted ID is not ObjectID.")
	}
	user.ID = id
	return user, nil
}
