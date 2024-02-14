package repository

import (
	"context"
	"crud/internal/redis/usecase"
	"log"

	"github.com/redis/go-redis/v9"
)

type repository struct {
	redisConnection *redis.Client
}

func NewRepository(conn *redis.Client) usecase.RepositoryInterfaces {
	return &repository{redisConnection: conn}
}

func (r *repository) CreateRedis(key, value string) error {
	if err := r.redisConnection.Set(context.Background(), key, value, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (r *repository) ReadRedis(key string) (string, error) {
	log.Println(key)
	value, err := r.redisConnection.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func (r *repository) UpdateRedis(key, value string) error {
	if err := r.redisConnection.Set(context.Background(), key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteRedis(key string) error {
	if err := r.redisConnection.Del(context.Background(), key).Err(); err != nil {
		return err
	}
	return nil
}
