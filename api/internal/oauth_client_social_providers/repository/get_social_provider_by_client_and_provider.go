package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/crypto"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/domain"
)

type getSocialProviderByClientAndProviderRepository struct {
	database types.Database
}

func (r *getSocialProviderByClientAndProviderRepository) Execute(ctx context.Context, clientID, provider string) (domain.SocialProvider, error) {
	var result domain.SocialProvider
	query := `SELECT id, client_id, provider, client_key, client_secret, enabled, scopes, created_at, updated_at
		FROM oauth_client_social_providers WHERE client_id = $1 AND provider = $2`
	err := r.database.QueryRowxContext(ctx, query, clientID, provider).StructScan(&result)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("GetSocialProviderByClientAndProviderRepository", err)
	}

	result.ClientKey, err = crypto.Decrypt(result.ClientKey)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("GetSocialProviderByClientAndProviderRepository", err)
	}
	result.ClientSecret, err = crypto.Decrypt(result.ClientSecret)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("GetSocialProviderByClientAndProviderRepository", err)
	}
	return result, nil
}

func NewGetSocialProviderByClientAndProviderRepository(database types.Database) contract.GetSocialProviderByClientAndProviderRepository {
	return &getSocialProviderByClientAndProviderRepository{database: database}
}
