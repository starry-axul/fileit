package client

import (
	"context"
	"github.com/starry-axul/fileit/internal/domain"
)

type (
	Service interface {
		Create(ctx context.Context, name string) (*domain.Client, error)
		GetAll(ctx context.Context, offset, limit int) ([]domain.Client, error)
		Count(ctx context.Context) (int, error)
	}
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, name string) (*domain.Client, error) {
	client := domain.Client{
		Name: name,
	}

	if err := s.repo.Create(ctx, &client); err != nil {
		return nil, err
	}

	return &client, nil
}

func (s service) GetAll(ctx context.Context, offset, limit int) ([]domain.Client, error) {

	clients, err := s.repo.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (s service) Count(ctx context.Context) (int, error) {

	count, err := s.repo.Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
