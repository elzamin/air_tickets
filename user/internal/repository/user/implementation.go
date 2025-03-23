package user

import (
	"github.com/elzamin/air_tickets/user/internal/entity/model"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func New(
	db *pgxpool.Pool,
) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) TouchTable(ctx context.Context) error {
	_, err := r.db.Exec(
		ctx,
		`
		CREATE TABLE IF NOT EXISTS userr(
			id text,
			first_name text,
			last_name text,
			PRIMARY KEY (id)
		)
		`,
	)

	return err
}

func (r *repository) Create(ctx context.Context, user model.User) error {
	_, err := r.db.Exec(
		ctx,
		"INSERT INTO userr (id, first_name, last_name) VALUES ($1, $2, $3)",
		user.Id,
		user.FirstName,
		user.LastName,
	)

	return err
}

func (r *repository) Get(ctx context.Context, id string) (model.User, error) {
	row := r.db.QueryRow(
		ctx,
		"SELECT id, first_name, last_name FROM userr WHERE id = $1",
		id,
	)

	var user model.User
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
