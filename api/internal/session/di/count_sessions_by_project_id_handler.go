package di

import (
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/booscaaa/locksmith/api/internal/session/contract"
	"github.com/booscaaa/locksmith/api/internal/session/handler"
	"github.com/booscaaa/locksmith/api/internal/session/repository"
	"github.com/booscaaa/locksmith/api/internal/session/usecase"
)

func NewCountSessionsByProjectIDHandler() contract.CountSessionsByProjectIDHandler {
	return handler.NewCountSessionsByProjectIDHandler(
		usecase.NewCountSessionsByProjectIDUseCase(
			repository.NewCountSessionsByProjectIDRepository(
				database.GetConnection(),
			),
		),
	)
}
