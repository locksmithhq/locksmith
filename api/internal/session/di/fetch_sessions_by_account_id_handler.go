package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/locksmithhq/locksmith/api/internal/session/handler"
	"github.com/locksmithhq/locksmith/api/internal/session/repository"
	"github.com/locksmithhq/locksmith/api/internal/session/usecase"
)

func NewFetchSessionsByAccountIDHandler() contract.FetchSessionsByAccountIDHandler {
	return handler.NewFetchSessionsByAccountIDHandler(
		usecase.NewFetchSessionsByAccountIDUseCase(
			repository.NewFetchSessionsByAccountIDRepository(
				database.GetConnection(),
			),
		),
	)
}
