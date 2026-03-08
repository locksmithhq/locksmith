package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/usecase"
)

func NewCreateLoginByClientIDHandler() contract.CreateLoginByClientIDHandler {
	return handler.NewCreateLoginByClientIDHandler(
		usecase.NewCreateLoginByClientIDUseCase(
			repository.NewCreateLoginByClientIDRepository(
				database.GetConnection(),
			),
		),
	)
}
