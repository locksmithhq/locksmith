package di

import (
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/handler"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/repository"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/usecase"
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
