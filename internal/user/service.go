package user

import (
	"context"
	"errors"
	"fmt"
	"go-auth/internal/auth"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *Repo
	jwtSecret string
}

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResult struct {
	Token string     `json:"token"`
	User  PublicUser `json:"user"`
}

func NewService(repo *Repo, jwtSecret string) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.TrimSpace(input.Password)

	if email == "" || password == "" {
		return AuthResult{}, errors.New("email and password are required")
	}

	if len(password) < 6 {
		return AuthResult{}, errors.New("password must be at least 6 characters long")
	}

	//check if user already exists
	_, err := s.repo.FindByEmail(ctx, email)
	if err == nil {
		return AuthResult{}, errors.New("Email is already registered! Please try with a different email.")
	}

	if !errors.Is(err, mongo.ErrNoDocuments) {
		return AuthResult{}, err
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResult{}, fmt.Errorf("Hashing of password failed: %w", err)
	}

	now := time.Now().UTC()

	user := User{
		Email:        email,
		PasswordHash: string(hashBytes),
		Role:         "user",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	created, err := s.repo.Create(ctx, user)
	if err != nil {
		return AuthResult{}, err
	}

	token, err := auth.CreateToken(s.jwtSecret, created.ID.Hex(), created.Role)

	if err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		Token: token,
		User:  ToPublic(created),
	}, nil
}
