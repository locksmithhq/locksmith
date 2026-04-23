package di

import (
	"github.com/locksmithhq/locksmith/api/internal/adapter/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/handler"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/repository"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/usecase"
)

func NewSocialCallbackHandler() contract.SocialCallbackHandler {
	return handler.NewSocialCallbackHandler(
		usecase.NewSocialCallbackUseCase(
			repository.NewGetSocialStateByNonceRepository(database.GetConnection()),
			repository.NewDeleteSocialStateRepository(database.GetConnection()),
			repository.NewGetClientByClientIDRepository(database.GetConnection()),
			repository.NewGetSocialProviderByClientRepository(database.GetConnection()),
			repository.NewGetAccountBySocialProviderRepository(database.GetConnection()),
			repository.NewGetAccountByEmailAndProjectRepository(database.GetConnection()),
			repository.NewCreateAccountSocialProviderRepository(database.GetConnection()),
			repository.NewCreateAccountRepository(database.GetConnection()),
			repository.NewCreateAuthCodeRepository(database.GetConnection()),
			repository.NewGetSignupByClientIDRepository(database.GetConnection()),
		),
	)
}
