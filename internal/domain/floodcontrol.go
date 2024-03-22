package domain

import (
	"context"
	"github.com/pkg/errors"
	"task/internal/config"
	"task/internal/database/repository"
)

var LimitError = errors.New("Request limit exceeded")

// FloodControl интерфейс, который нужно реализовать.
// Рекомендуем создать директорию-пакет, в которой будет находиться реализация.
type FloodControl interface {
	// Check возвращает false если достигнут лимит максимально разрешенного
	// кол-ва запросов согласно заданным правилам флуд контроля.
	Check(ctx context.Context, userID int64) (bool, error)
}

// FloodController - имплементация FloodControl
type FloodController struct {
	maxRequests int
	repository  repository.RamRepository
}

func NewFloodController(cfg config.AppConfig, db repository.RamRepository) *FloodController {
	return &FloodController{maxRequests: cfg.MaxRequests, repository: db}
}

func (f *FloodController) Check(ctx context.Context, userID int64) (bool, error) {
	userQuery, err := f.repository.GetFirstByUserId(ctx, userID)

	if err != nil {
		if err = f.repository.Save(ctx, repository.UserQuery{UserId: userID, Tokens: 0}); err != nil {
			return false, err
		}
	} else {
		if userQuery.Tokens > f.maxRequests {
			return false, LimitError
		} else {
			if err = f.repository.Save(ctx, userQuery); err != nil {
				return false, err
			}
		}
	}

	return true, nil
}
