package user

import (
	"airtickets/internal/entity/model"
	"context"
)

type Repository interface {
	TouchTable(ctx context.Context) error
	Create(ctx context.Context, user model.User) error
	Get(ctx context.Context, id string) (model.User, error)
}