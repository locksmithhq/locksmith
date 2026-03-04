package di

import (
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/handler"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/repository"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/usecase"
)

func NewGetClientByIDAndProjectIDHandler() contract.GetClientByIDAndProjectIDHandler {
	return handler.NewGetClientByIDAndProjectIDHandler(
		usecase.NewGetClientByIDAndProjectIDUseCase(
			repository.NewGetClientByIDAndProjectIDRepository(database.GetConnection()),
		),
	)
}
