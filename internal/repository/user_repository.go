package repository

import (
	"context"

	"github.com/Caixetadev/accept-interfaces-return-structs/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	database *pgxpool.Pool
}

func NewUserRepository(database *pgxpool.Pool) *userRepository {
	return &userRepository{database: database}
}

func (ur *userRepository) Save(ctx context.Context, user *domain.User) error {
	_, err := ur.database.Exec(ctx, "INSERT INTO users(id, name, lastname) VALUES($1, $2, $3)", user.ID, user.Name, user.LastName)

	if err != nil {
		return err
	}

	return nil
}
