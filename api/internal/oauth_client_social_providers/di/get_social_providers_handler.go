package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/usecase"
)

func NewGetSocialProvidersHandler() contract.GetSocialProvidersByClientIDHandler {
	return handler.NewGetSocialProvidersByClientIDHandler(
		usecase.NewGetSocialProvidersByClientIDUseCase(
			repository.NewGetSocialProvidersByClientIDRepository(database.GetConnection()),
		),
	)
}
