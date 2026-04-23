package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/domain"
)

type GetSocialProvidersByClientIDRepository interface {
	Execute(context.Context, string) ([]domain.SocialProvider, error)
}

type GetSocialProviderByClientAndProviderRepository interface {
	Execute(ctx context.Context, clientID, provider string) (domain.SocialProvider, error)
}

type UpsertSocialProviderRepository interface {
	Execute(context.Context, domain.SocialProvider) (domain.SocialProvider, error)
}
