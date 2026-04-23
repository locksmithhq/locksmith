package di

import (
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/handler"
	"github.com/locksmithhq/locksmith/api/internal/account/repository"
	"github.com/locksmithhq/locksmith/api/internal/account/usecase"
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
)

func NewFetchAccountsByProjectIDHandler() contract.FetchAccountsByProjectIDHandler {
	return handler.NewFetchAccountsByProjectIDHandler(
		usecase.NewFetchAccountsByProjectIDUseCase(
			repository.NewFetchAccountsByProjectIDRepository(database.GetConnection()),
			repository.NewFetchSocialProvidersByAccountIDsRepository(database.GetConnection()),
		),
	)
}
