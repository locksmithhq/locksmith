package di

import (
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/handler"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/repository"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/usecase"
)

func NewFetchClientsByProjectIDHandler() contract.FetchClientsByProjectIDHandler {
	return handler.NewFetchClientsByProjectIDHandler(
		usecase.NewFetchClientsByProjectIDUseCase(
			repository.NewFetchClientsByProjectIDRepository(database.GetConnection()),
		),
	)
}
