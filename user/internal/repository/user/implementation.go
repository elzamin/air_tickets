package user

import (
	"github.com/elzamin/air_tickets/user/internal/entity"
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

func (r *repository) Create(ctx context.Context, user entity.User) error {
	_, err := r.db.Exec(
		ctx,
		"INSERT INTO userr (id, first_name, last_name) VALUES ($1, $2, $3)",
		user.Id,
		user.FirstName,
		user.LastName,
	)

	return err
}

func (r *repository) Get(ctx context.Context, id string) (entity.User, error) {
	row := r.db.QueryRow(
		ctx,
		"SELECT id, first_name, last_name FROM userr WHERE id = $1",
		id,
	)

	var user entity.User
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *repository) Update(ctx context.Context, user entity.User) error {
	_, err := r.db.Exec(
		ctx,
		"UPDATE userr SET first_name = $1, last_name = $2 WHERE id = $3",
		user.FirstName,
		user.LastName,
		user.Id,
	)

	return err
}

func (r *repository) Delete(ctx context.Context, id string) error{
	_, err := r.db.Exec(
		ctx,
		"DELETE FROM userr WHERE id = $1",
		id,
	)

	return err
}