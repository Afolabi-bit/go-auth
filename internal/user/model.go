package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	PasswordHash string             `json:"-" bson:"passwordHash"`
	Role         string             `json:"role" bson:"role"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
}

type PublicUser struct {
	ID        string    `json:"id" bson:"_id"`
	Email     string    `json:"email" bson:"email"`
	Role      string    `json:"role" bson:"role"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

func ToPublic(u User) PublicUser {
	return PublicUser{
		ID:        u.ID.Hex(),
		Email:     u.Email,
		Role:      u.Role,
		UpdatedAt: u.UpdatedAt,
		CreatedAt: u.CreatedAt,
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
