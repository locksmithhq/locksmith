package contract

import (
	"context"

	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/types/input"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/types/output"
)

type GetSocialProvidersByClientIDUseCase interface {
	Execute(context.Context, string) ([]output.SocialProvider, error)
}

type UpsertSocialProviderUseCase interface {
	Execute(context.Context, string, string, input.SocialProvider) (output.SocialProvider, error)
}
