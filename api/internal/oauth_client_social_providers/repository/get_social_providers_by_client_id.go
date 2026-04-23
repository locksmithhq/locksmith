package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/crypto"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/domain"
)

type getSocialProvidersByClientIDRepository struct {
	database types.Database
}

func (r *getSocialProvidersByClientIDRepository) Execute(ctx context.Context, clientID string) ([]domain.SocialProvider, error) {
	var providers []domain.SocialProvider

	query := `SELECT id, client_id, provider, client_key, client_secret, enabled, scopes, created_at, updated_at
		FROM oauth_client_social_providers WHERE client_id = $1`

	err := r.database.SelectContext(ctx, &providers, query, clientID)
	if err != nil {
		return nil, stackerror.NewRepositoryError("GetSocialProvidersByClientIDRepository", err)
	}

	for i := range providers {
		providers[i].ClientKey, err = crypto.Decrypt(providers[i].ClientKey)
		if err != nil {
			return nil, stackerror.NewRepositoryError("GetSocialProvidersByClientIDRepository", err)
		}
		providers[i].ClientSecret, err = crypto.Decrypt(providers[i].ClientSecret)
		if err != nil {
			return nil, stackerror.NewRepositoryError("GetSocialProvidersByClientIDRepository", err)
		}
	}
	return providers, nil
}

func NewGetSocialProvidersByClientIDRepository(database types.Database) contract.GetSocialProvidersByClientIDRepository {
	return &getSocialProvidersByClientIDRepository{database: database}
}
