package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/crypto"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/domain"
)

type upsertSocialProviderRepository struct {
	database types.Database
}

func (r *upsertSocialProviderRepository) Execute(ctx context.Context, provider domain.SocialProvider) (domain.SocialProvider, error) {
	var result domain.SocialProvider

	query := `INSERT INTO oauth_client_social_providers (client_id, provider, client_key, client_secret, enabled, scopes, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		ON CONFLICT (client_id, provider) DO UPDATE SET
			client_key = EXCLUDED.client_key,
			client_secret = EXCLUDED.client_secret,
			enabled = EXCLUDED.enabled,
			scopes = EXCLUDED.scopes,
			updated_at = NOW()
		RETURNING id, client_id, provider, client_key, client_secret, enabled, scopes, created_at, updated_at`

	encryptedKey, err := crypto.Encrypt(provider.ClientKey)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("UpsertSocialProviderRepository", err)
	}
	encryptedSecret, err := crypto.Encrypt(provider.ClientSecret)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("UpsertSocialProviderRepository", err)
	}

	err = r.database.QueryRowxContext(ctx, query,
		provider.ClientID,
		provider.Provider,
		encryptedKey,
		encryptedSecret,
		provider.Enabled,
		provider.Scopes,
	).StructScan(&result)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("UpsertSocialProviderRepository", err)
	}

	result.ClientKey, err = crypto.Decrypt(result.ClientKey)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("UpsertSocialProviderRepository", err)
	}
	result.ClientSecret, err = crypto.Decrypt(result.ClientSecret)
	if err != nil {
		return domain.SocialProvider{}, stackerror.NewRepositoryError("UpsertSocialProviderRepository", err)
	}
	return result, nil
}

func NewUpsertSocialProviderRepository(database types.Database) contract.UpsertSocialProviderRepository {
	return &upsertSocialProviderRepository{database: database}
}
