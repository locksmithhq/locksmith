package di

import (
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/handler"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/repository"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/usecase"
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
