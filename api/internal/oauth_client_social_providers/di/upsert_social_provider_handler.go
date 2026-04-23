package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/usecase"
)

func NewUpsertSocialProviderHandler() contract.UpsertSocialProviderHandler {
	conn := database.GetConnection()
	return handler.NewUpsertSocialProviderHandler(
		usecase.NewUpsertSocialProviderUseCase(
			repository.NewUpsertSocialProviderRepository(conn),
			repository.NewGetSocialProviderByClientAndProviderRepository(conn),
		),
	)
}
