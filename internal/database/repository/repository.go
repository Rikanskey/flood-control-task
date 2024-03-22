package repository

import "context"

// RamRepository интерфейс, декларирующий взаимодействие с хранилищем данных
type RamRepository interface {
	GetFirstByUserId(ctx context.Context, userId int64) (UserQuery, error)
	Save(ctx context.Context, query UserQuery) error
}
