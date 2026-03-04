package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/cmd"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/usecase"
)

func NewConfigDefaultClientCMD() contract.ConfigDefaultClientCMD {
	return cmd.NewConfigDefaultClientCMD(
		usecase.NewCreateClientUseCase(
			repository.NewCreateClientRepository(database.GetConnection()),
			repository.NewGetClientByProjectIDAndClientIDRepository(database.GetConnection()),
		),
	)
}
