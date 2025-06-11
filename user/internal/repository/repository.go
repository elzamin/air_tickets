package repository

import (
	"context"
	"errors"

	"github.com/elzamin/air_tickets/user/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func New(
	pool *pgxpool.Pool,
) *repository {
	return &repository{
		db: pool,
	}
}

func (r *repository) TouchTable(ctx context.Context) error {
	_, err := r.db.Exec(
		ctx,
		`
		CREATE TABLE IF NOT EXISTS usertable(
			id text,
			name text,
			age integer,
			address text,
			work text,
			PRIMARY KEY (id)
		)
		`,
	)
	
	return err
}

func (r *repository) Create(ctx context.Context, user entity.User) error {
	_, err := r.db.Exec(
		ctx,
		"INSERT INTO usertable (id, name, age, address, work) VALUES ($1, $2, $3, $4, $5)",
		user.Id,
		user.Name,
		user.Age,
		user.Address,
		user.Work,
	)

	return err
}

func (r *repository) Get(ctx context.Context, id string) (entity.User, error) {
	row := r.db.QueryRow(
		ctx,
		"SELECT id, name, age, address, work FROM usertable WHERE id = $1",
		id,
	)

	var user entity.User
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Age,
		&user.Address,
		&user.Work)
	if err != nil {
		return entity.User{}, errors.New("DB.Scan fail: ")
	}

	return user, nil
}

func (r *repository) GetAll(ctx context.Context) ([]entity.User, error) {
	rows, err := r.db.Query(
		ctx,
		"SELECT id, name, age, address, work FROM usertable",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Age,
			&user.Address,
			&user.Work)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repository) Update(ctx context.Context, user entity.User) error {
	_, err := r.db.Exec(
		ctx,
		"UPDATE usertable SET name = $1, age = $2, address = $3, work = $4 WHERE id = $5",
		user.Name,
		user.Age,
		user.Address,
		user.Work,
		user.Id,
	)

	return err
}

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(
		ctx,
		"DELETE FROM usertable WHERE id = $1",
		id,
	)

	return err
}
