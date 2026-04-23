package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/crypto"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getSocialProviderByClientRepository struct {
	database types.Database
}

func (r *getSocialProviderByClientRepository) Execute(ctx context.Context, clientID string, provider string) (domain.SocialProvider, error) {
	var result domain.SocialProvider

	query := `SELECT id, client_id, provider, client_key, client_secret, enabled, scopes
		FROM oauth_client_social_providers WHERE client_id = $1 AND provider = $2`

	err := r.database.QueryRowxContext(ctx, query, clientID, provider).StructScan(&result)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("GetSocialProviderByClientRepository", err)
	}

	result.ClientKey, err = crypto.Decrypt(result.ClientKey)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("GetSocialProviderByClientRepository", err)
	}
	result.ClientSecret, err = crypto.Decrypt(result.ClientSecret)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("GetSocialProviderByClientRepository", err)
	}
	return result, nil
}

func NewGetSocialProviderByClientRepository(database types.Database) contract.GetSocialProviderByClientRepository {
	return &getSocialProviderByClientRepository{database: database}
}
