package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/locksmithhq/locksmith/api/internal/session/handler"
	"github.com/locksmithhq/locksmith/api/internal/session/repository"
	"github.com/locksmithhq/locksmith/api/internal/session/usecase"
)

func NewCountSessionsByAccountIDHandler() contract.CountSessionsByAccountIDHandler {
	return handler.NewCountSessionsByAccountIDHandler(
		usecase.NewCountSessionsByAccountIDUseCase(
			repository.NewCountSessionsByAccountIDRepository(
				database.GetConnection(),
			),
		),
	)
}
