package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type createAccountSocialProviderRepository struct {
	database types.Database
}

func (r *createAccountSocialProviderRepository) Execute(ctx context.Context, asp domain.AccountSocialProvider) error {
	query := `INSERT INTO account_social_providers (account_id, provider, provider_user_id, email)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (provider, provider_user_id) DO NOTHING`

	_, err := r.database.ExecContext(ctx, query,
		asp.AccountID,
		asp.Provider,
		asp.ProviderUserID,
		asp.Email,
	)
	if err != nil {
		return stackerror.NewRepositoryError("CreateAccountSocialProviderRepository", err)
	}
	return nil
}

func NewCreateAccountSocialProviderRepository(database types.Database) contract.CreateAccountSocialProviderRepository {
	return &createAccountSocialProviderRepository{database: database}
}
