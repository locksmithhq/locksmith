package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getAccountBySocialProviderRepository struct {
	database types.Database
}

func (r *getAccountBySocialProviderRepository) Execute(ctx context.Context, provider string, providerUserID string) (domain.Account, error) {
	var account domain.Account

	query := `SELECT a.id, a.project_id, a.name, a.email, a.username, a.password, a.created_at, a.updated_at, a.deleted_at, a.must_change_password
		FROM accounts a
		JOIN account_social_providers asp ON asp.account_id = a.id
		WHERE asp.provider = $1 AND asp.provider_user_id = $2 AND a.deleted_at IS NULL`

	err := r.database.QueryRowxContext(ctx, query, provider, providerUserID).StructScan(&account)
	if err != nil {
		return domain.Account{}, stackerror.NewRepositoryError("GetAccountBySocialProviderRepository", err)
	}
	return account, nil
}

func NewGetAccountBySocialProviderRepository(database types.Database) contract.GetAccountBySocialProviderRepository {
	return &getAccountBySocialProviderRepository{database: database}
}
