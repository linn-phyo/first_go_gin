//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"gorm.io/gorm"

	handler "github.com/linn-phyo/go_gin_clean_architecture/src/api/handler"
	repository "github.com/linn-phyo/go_gin_clean_architecture/src/repository"
	usecase "github.com/linn-phyo/go_gin_clean_architecture/src/usecase"
)

func InitializeUsersAPI(DB *gorm.DB) (*handler.UserHandler, error) {
	wire.Build(repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler)

	return &handler.UserHandler{}, nil
}
