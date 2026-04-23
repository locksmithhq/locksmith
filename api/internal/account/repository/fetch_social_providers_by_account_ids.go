package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/lib/pq"
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
)

type fetchSocialProvidersByAccountIDsRepository struct {
	database types.Database
}

func (r *fetchSocialProvidersByAccountIDsRepository) Execute(ctx context.Context, accountIDs []string) (map[string][]string, error) {
	result := make(map[string][]string)
	if len(accountIDs) == 0 {
		return result, nil
	}

	type row struct {
		AccountID string `db:"account_id"`
		Provider  string `db:"provider"`
	}
	var rows []row

	query := `SELECT account_id, provider FROM account_social_providers WHERE account_id = ANY($1)`
	if err := r.database.SelectContext(ctx, &rows, query, pq.Array(accountIDs)); err != nil {
		return nil, stackerror.NewRepositoryError("FetchSocialProvidersByAccountIDsRepository", err)
	}

	for _, row := range rows {
		result[row.AccountID] = append(result[row.AccountID], row.Provider)
	}
	return result, nil
}

func NewFetchSocialProvidersByAccountIDsRepository(database types.Database) contract.FetchSocialProvidersByAccountIDsRepository {
	return &fetchSocialProvidersByAccountIDsRepository{database: database}
}
