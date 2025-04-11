package repository

import (
	"context"
	"github.com/elzamin/air_tickets/user/internal/entity"
)

type Repository interface {
	TouchTable(ctx context.Context) error
	Create(ctx context.Context, user entity.User) error
	Get(ctx context.Context, id string) (entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, id string) error
}
