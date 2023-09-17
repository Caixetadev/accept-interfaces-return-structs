package main

import (
	"context"

	"github.com/Caixetadev/accept-interfaces-return-structs/internal/database"
	"github.com/Caixetadev/accept-interfaces-return-structs/internal/domain"
	"github.com/Caixetadev/accept-interfaces-return-structs/internal/repository"
	"github.com/Caixetadev/accept-interfaces-return-structs/internal/services"
)

func main() {
	db := database.Connect()

	defer db.Close()

	repository := repository.NewUserRepository(db)
	services := services.NewUserService(repository)

	input := &domain.User{
		Name:     "Rafael",
		LastName: "Caixeta",
	}

	_, err := services.CreateUser(context.Background(), input)
	if err != nil {
		panic(err)
	}
}
