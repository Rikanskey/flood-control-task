package db

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"task/internal/config"
	"task/internal/database/repository"
	"time"
)

// RedisRepository имплементация RamRepository ориентированая на

type RedisRepository struct {
	refreshTime time.Duration
	client      *redis.Client
}

func NewRedisRepository(rc config.RedisConfig) (*RedisRepository, error) {
	redRep := RedisRepository{refreshTime: rc.RefreshTime}
	client := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Password: rc.Password,
		DB:       rc.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	redRep.client = client
	return &redRep, nil
}

func (r *RedisRepository) GetFirstByUserId(ctx context.Context, userId int64) (repository.UserQuery, error) {
	var query repository.UserQuery

	data, err := r.client.Get(ctx, string(userId)).Bytes()
	if err != nil {
		return repository.UserQuery{}, err
	}

	err = json.Unmarshal(data, &query)
	if err != nil {
		return repository.UserQuery{}, err
	}

	return query, nil
}

func (r *RedisRepository) Save(ctx context.Context, userQuery repository.UserQuery) error {
	jsonString, err := json.Marshal(repository.UserQuery{
		Tokens: userQuery.Tokens + 1,
		Time:   time.Now(),
	})
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, string(userQuery.UserId), jsonString, r.refreshTime).Err()

	if err != nil {
		return err
	}

	return nil
}
