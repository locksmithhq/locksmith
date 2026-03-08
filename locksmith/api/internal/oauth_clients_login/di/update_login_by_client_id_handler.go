package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/usecase"
)

func NewUpdateLoginByClientIDHandler() contract.UpdateLoginByClientIDHandler {
	return handler.NewUpdateLoginByClientIDHandler(
		usecase.NewUpdateLoginByClientIDUseCase(
			repository.NewUpdateLoginByClientIDRepository(
				database.GetConnection(),
			),
		),
	)
}
