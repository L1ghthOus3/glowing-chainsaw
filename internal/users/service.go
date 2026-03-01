package users

import (
	"context"
	"errors"
	"strings"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, in CreateUserInput) (User, error) {
	in.Email = strings.TrimSpace(in.Email)
	in.Name = strings.TrimSpace(in.Name)

	if in.Email == "" || !strings.Contains(in.Email, "@") {
		return User{}, errors.New("invalid email")
	}
	if in.Name == "" {
		return User{}, errors.New("name is required")
	}

	return s.repo.Create(ctx, in)
}

func (s *Service) GetByID(ctx context.Context, id int) (User, error) {
	if id <= 0 {
		return User{}, errors.New("invalid id")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context, limit int) ([]User, error) {
	return s.repo.List(ctx, limit)
}
