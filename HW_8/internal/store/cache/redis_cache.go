package cache

import (
	"HW_8/internal/models"
	"context"
)

type Cache interface {
	User() UserCache
}

type UserCache interface {
	Set(ctx context.Context, key string, data *models.User)
	Get(ctx context.Context, key string) *models.User
}
