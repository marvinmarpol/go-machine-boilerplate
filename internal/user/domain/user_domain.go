package domain

import (
	"errors"

	"github.com/gofrs/uuid"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewUser(email, name string) (*User, error) {
	if email == "" || name == "" {
		return nil, errors.New("email and name must be provided")
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    uuid.String(),
		Email: email,
		Name:  name,
	}, nil
}
