package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type getEnabledSocialProvidersByClientRepository struct {
	database types.Database
}

func (r *getEnabledSocialProvidersByClientRepository) Execute(ctx context.Context, clientID string) ([]string, error) {
	var providers []string

	query := `SELECT provider FROM oauth_client_social_providers
		WHERE client_id = $1 AND enabled = true`

	err := r.database.SelectContext(ctx, &providers, query, clientID)
	if err != nil {
		return nil, stackerror.NewRepositoryError("GetEnabledSocialProvidersByClientRepository", err)
	}
	return providers, nil
}

func NewGetEnabledSocialProvidersByClientRepository(database types.Database) contract.GetEnabledSocialProvidersByClientRepository {
	return &getEnabledSocialProvidersByClientRepository{database: database}
}
