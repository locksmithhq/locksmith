package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/usecase"
)

func NewFetchClientsByProjectIDHandler() contract.FetchClientsByProjectIDHandler {
	return handler.NewFetchClientsByProjectIDHandler(
		usecase.NewFetchClientsByProjectIDUseCase(
			repository.NewFetchClientsByProjectIDRepository(database.GetConnection()),
		),
	)
}
