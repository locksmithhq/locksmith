package di

import (
	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/account/handler"
	"github.com/booscaaa/locksmith/api/internal/account/repository"
	"github.com/booscaaa/locksmith/api/internal/account/usecase"
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
)

func NewCountAccountsByProjectIDHandler() contract.CountAccountsByProjectIDHandler {
	return handler.NewCountAccountsByProjectIDHandler(
		usecase.NewCountAccountsByProjectIDUseCase(
			repository.NewCountAccountsByProjectIDRepository(
				database.GetConnection(),
			),
		),
	)
}
