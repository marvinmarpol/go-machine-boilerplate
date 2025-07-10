package adapter

import (
	"database/sql"
	"go-machine-boilerplate/internal/user/domain"
)

type UserAdapterPostgres struct {
	DB *sql.DB
}

func NewUserAdapterPostgres(db *sql.DB) *UserAdapterPostgres {
	return &UserAdapterPostgres{DB: db}
}

func (r *UserAdapterPostgres) Save(user *domain.User) (string, error) {
	row := r.DB.QueryRow("INSERT INTO public.users (id, email, name) VALUES ($1, $2, $3) RETURNING id", user.ID, user.Email, user.Name)

	var u domain.User
	if err := row.Scan(&u.ID); err != nil {
		return "", err
	}
	return u.ID, nil
}

func (r *UserAdapterPostgres) FindById(id string) (*domain.User, error) {
	row := r.DB.QueryRow("SELECT id, email, name FROM public.users WHERE id = $1", id)

	var u domain.User
	if err := row.Scan(&u.ID, &u.Email, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserAdapterPostgres) FindByEmail(email string) (*domain.User, error) {
	row := r.DB.QueryRow("SELECT id, email, name FROM public.users WHERE email = $1", email)

	var u domain.User
	if err := row.Scan(&u.ID, &u.Email, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}
