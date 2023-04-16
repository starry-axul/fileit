package client

import (
	"context"
	"github.com/starry-axul/fileit/internal/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, client *domain.Client) error
	GetAll(ctx context.Context, offset, limit int) ([]domain.Client, error)
	Update(ctx context.Context, id int) error
	Count(ctx context.Context) (int, error)
}

// repo persists users in database
type repo struct {
	db *gorm.DB
}

// NewRepository creates a new users repository
func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) Create(ctx context.Context, client *domain.Client) error {
	return r.db.WithContext(ctx).Create(client).Error
}

func (r *repo) GetAll(ctx context.Context, offset, limit int) ([]domain.Client, error) {
	var c []domain.Client

	tx := r.db.WithContext(ctx).Model(&c)
	tx = tx.Limit(limit).Offset(offset)

	if err := tx.Order("created_at desc").Find(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (r *repo) Update(ctx context.Context, id int) error {
	return nil
}

func (r *repo) Count(ctx context.Context) (int, error) {

	var count int64

	tx := r.db.WithContext(ctx).Model(domain.Client{})

	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
