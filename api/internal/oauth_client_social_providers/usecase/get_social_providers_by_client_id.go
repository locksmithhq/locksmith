package usecase

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/types/output"
)

type getSocialProvidersByClientIDUseCase struct {
	getSocialProvidersByClientIDRepository contract.GetSocialProvidersByClientIDRepository
}

func (u *getSocialProvidersByClientIDUseCase) Execute(ctx context.Context, clientID string) ([]output.SocialProvider, error) {
	providers, err := u.getSocialProvidersByClientIDRepository.Execute(ctx, clientID)
	if err != nil {
		return nil, stackerror.NewUseCaseError("GetSocialProvidersByClientIDUseCase", err)
	}
	return output.NewSocialProvidersFromDomain(providers), nil
}

func NewGetSocialProvidersByClientIDUseCase(
	getSocialProvidersByClientIDRepository contract.GetSocialProvidersByClientIDRepository,
) contract.GetSocialProvidersByClientIDUseCase {
	return &getSocialProvidersByClientIDUseCase{
		getSocialProvidersByClientIDRepository: getSocialProvidersByClientIDRepository,
	}
}
