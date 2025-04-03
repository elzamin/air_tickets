package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/elzamin/air_tickets/user/internal/infrastructure/model"
	"github.com/pkg/errors"
)

func NewPostgres(cfg model.Postgres) (*pgxpool.Pool, error) {
	conn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Database,
		cfg.Password,
	)

	ctx := context.Background()
	db, err := pgxpool.New(ctx, conn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pgx.Connect")
	}

	if err = db.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to ping a db")
	}

	return db, nil
}
