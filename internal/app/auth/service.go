package auth

import (
	"context"

	"hackathon-dsm-server/internal/pkg/utils/errors"
	"hackathon-dsm-server/internal/pkg/utils/jwt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, req *RegisterRequest) (*AuthData, error)
	Login(ctx context.Context, req *LoginRequest) (*AuthData, error)
}

type AuthService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, req *RegisterRequest) (*AuthData, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.InternalServerError("Failed to hash password")
	}

	user := &User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Nickname: req.Nickname,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	token, err := jwt.GenerateAccessToken(user.UserID)
	if err != nil {
		return nil, errors.InternalServerError("Failed to generate token")
	}

	return &AuthData{
		UserID: user.UserID,
		Token:  token,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*AuthData, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.Unauthorized("Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.Unauthorized("Invalid email or password")
	}

	token, err := jwt.GenerateAccessToken(user.UserID)
	if err != nil {
		return nil, errors.InternalServerError("Failed to generate token")
	}

	return &AuthData{
		UserID:   user.UserID,
		Nickname: user.Nickname,
		Token:    token,
	}, nil
}