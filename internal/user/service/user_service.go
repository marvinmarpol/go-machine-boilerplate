package service

import (
	"database/sql"
	"go-machine-boilerplate/internal/user/domain"
	"strings"
)

type UserRepository interface {
	Save(user *domain.User) (string, error)
	FindById(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (s *UserService) Create(email, name string) (string, error) {

	existing, err := s.repository.FindByEmail(email)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if existing != nil && strings.TrimSpace(existing.ID) != "" {
		return existing.ID, nil
	}

	user, err := domain.NewUser(email, name)
	if err != nil {
		return "", err
	}

	return s.repository.Save(user)
}

func (s *UserService) Get(id string) (*domain.User, error) {
	return s.repository.FindById(id)
}
