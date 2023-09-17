package repository

import (
	"context"

	"github.com/Caixetadev/accept-interfaces-return-structs/internal/domain"
)

type MockUserRepository struct {
	SaveCalled bool
}

func (m *MockUserRepository) Save(ctx context.Context, user *domain.User) error {
	m.SaveCalled = true
	return nil
}
