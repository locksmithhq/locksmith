package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/locksmithhq/locksmith/api/internal/session/handler"
	"github.com/locksmithhq/locksmith/api/internal/session/repository"
	"github.com/locksmithhq/locksmith/api/internal/session/usecase"
)

func NewRevokeSessionHandler() contract.RevokeSessionHandler {
	return handler.NewRevokeSessionHandler(
		usecase.NewRevokeSessionUseCase(
			repository.NewRevokeSessionRepository(
				database.GetConnection(),
			),
		),
	)
}
