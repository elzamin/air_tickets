package user

import (
	"github.com/elzamin/air_tickets/user/internal/entity/model"
	"context"
)

type Repository interface {
	TouchTable(ctx context.Context) error
	Create(ctx context.Context, user model.User) error
	Get(ctx context.Context, id string) (model.User, error)
}