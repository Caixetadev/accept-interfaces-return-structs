package services

import (
	"context"
	"testing"

	"github.com/Caixetadev/accept-interfaces-return-structs/internal/domain"
	"github.com/Caixetadev/accept-interfaces-return-structs/internal/repository"
)

func TestCreateUser(t *testing.T) {
	mockRepository := &repository.MockUserRepository{}

	service := NewUserService(mockRepository)

	user := &domain.User{
		Name:     "Rafael",
		LastName: "Caixeta",
	}

	userCreated, err := service.CreateUser(context.Background(), user)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if !mockRepository.SaveCalled {
		t.Error("Expected Save to be called on UserRepository")
	}

	if userCreated.ID == "" {
		t.Error("Expected user ID to be assigned, but got an empty ID")
	}
}
