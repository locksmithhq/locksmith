package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/usecase"
)

func NewGetClientByIDAndProjectIDHandler() contract.GetClientByIDAndProjectIDHandler {
	return handler.NewGetClientByIDAndProjectIDHandler(
		usecase.NewGetClientByIDAndProjectIDUseCase(
			repository.NewGetClientByIDAndProjectIDRepository(database.GetConnection()),
		),
	)
}
