package cache

import (
	"HW_8/internal/models"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserRedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewUserCache(host string, db int, expires time.Duration) UserCache {
	return &UserRedisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (a UserRedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     a.host,
		Password: "",
		DB:       a.db,
	})
}

func (u UserRedisCache) Set(ctx context.Context, key string, value *models.User) {
	client := u.getClient()
	user, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	_, err = client.Set(ctx, key, user, u.expires*time.Second).Result()
	if err != nil {
		return
	}
}

func (u UserRedisCache) Get(ctx context.Context, key string) *models.User {
	client := u.getClient()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	user := new(models.User)
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		panic(err)
	}
	return user
}
