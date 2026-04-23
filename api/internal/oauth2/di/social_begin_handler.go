package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/usecase"
)

func NewSocialBeginHandler() contract.SocialBeginHandler {
	return handler.NewSocialBeginHandler(
		usecase.NewSocialBeginUseCase(
			repository.NewGetClientByClientIDRepository(database.GetConnection()),
			repository.NewGetSocialProviderByClientRepository(database.GetConnection()),
			repository.NewCreateSocialStateRepository(database.GetConnection()),
		),
	)
}
