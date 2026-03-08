package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/locksmithhq/locksmith/api/internal/session/handler"
	"github.com/locksmithhq/locksmith/api/internal/session/repository"
	"github.com/locksmithhq/locksmith/api/internal/session/usecase"
)

func NewCountRefreshTokensByAccountIDHandler() contract.CountRefreshTokensByAccountIDHandler {
	return handler.NewCountRefreshTokensByAccountIDHandler(
		usecase.NewCountRefreshTokensByAccountIDUseCase(
			repository.NewCountRefreshTokensByAccountIDRepository(
				database.GetConnection(),
			),
		),
	)
}
