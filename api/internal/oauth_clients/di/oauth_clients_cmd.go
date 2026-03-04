package di

import (
	"github.com/booscaaa/locksmith/api/internal/adapter/database"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/cmd"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/repository"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/usecase"
)

func NewConfigDefaultClientCMD() contract.ConfigDefaultClientCMD {
	return cmd.NewConfigDefaultClientCMD(
		usecase.NewCreateClientUseCase(
			repository.NewCreateClientRepository(database.GetConnection()),
			repository.NewGetClientByProjectIDAndClientIDRepository(database.GetConnection()),
		),
	)
}
