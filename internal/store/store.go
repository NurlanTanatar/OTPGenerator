package store

import (
	"QRCodeGenerator/internal/models"
	"context"
)

type Store interface {
	Create(ctx context.Context, user *models.User) error
	All(ctx context.Context) ([]*models.User, error)
	ByID(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}
