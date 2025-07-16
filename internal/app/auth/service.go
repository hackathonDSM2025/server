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
	CheckUsername(ctx context.Context, req *CheckUsernameRequest) (*CheckUsernameData, error)
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
		Username: req.Username,
		Password: string(hashedPassword),
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
	user, err := s.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.Unauthorized("Invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.Unauthorized("Invalid username or password")
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

func (s *AuthService) CheckUsername(ctx context.Context, req *CheckUsernameRequest) (*CheckUsernameData, error) {
	_, err := s.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if customErr, ok := err.(*errors.CustomError); ok && customErr.Status == 404 {
			return &CheckUsernameData{
				Available: true,
				Message:   "Username is available",
			}, nil
		}
		return nil, err
	}

	return &CheckUsernameData{
		Available: false,
		Message:   "Username is already taken",
	}, nil
}