package service

import (
	"context"
	"errors"

	"github.com/elzamin/air_tickets/user/internal/entity"
)

type Service struct {
	userRepository UserRepository
}

func New(
	repo UserRepository,
) *Service {
	return &Service{
		userRepository: repo,
	}
}

type UserRepository interface {
	TouchTable(ctx context.Context) error
	Create(ctx context.Context, user entity.User) error
	Get(ctx context.Context, id string) (entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, id string) error
}

func (s *Service) Create(ctx context.Context, user entity.User) error {
	_, err := s.userRepository.Get(ctx, user.Id)
	if err == nil {
		return errors.New("User with ID: '" + user.Id + "' is exist")
	}
	return s.userRepository.Create(ctx, user)
}

func (s *Service) Get(ctx context.Context, id string) (entity.User, error) {
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return user, errors.New("User with ID: '" + id + "' is not exist")
	}

	return s.userRepository.Get(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]entity.User, error) {
	return s.userRepository.GetAll(ctx)
}

func (s *Service) Update(ctx context.Context, user entity.User) error {
	_, err := s.userRepository.Get(ctx, user.Id)
	if err != nil {
		return errors.New("User with ID: '" + user.Id + "' is not exist")
	}
	return s.userRepository.Update(ctx, user)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	_, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return errors.New("User with ID: '" + id + "' is not exist")
	}
	return s.userRepository.Delete(ctx, id)
}
