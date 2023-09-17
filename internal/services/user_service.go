package services

import (
	"context"
	"errors"

	"github.com/Caixetadev/accept-interfaces-return-structs/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	user = &domain.User{
		ID:       uuid.New().String(),
		Name:     user.Name,
		LastName: user.LastName,
	}

	err := us.userRepository.Save(ctx, user)
	if err != nil {
		return nil, errors.New("We encountered an error while processing your request.")
	}

	return user, nil
}
