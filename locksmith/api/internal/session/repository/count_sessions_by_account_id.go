package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type countSessionsByAccountIDRepository struct {
	database types.Database
}

func (r *countSessionsByAccountIDRepository) Execute(ctx context.Context, projectID, accountID string) (int64, error) {
	query := `
		SELECT COUNT(*)
		FROM user_sessions us
		JOIN oauth_clients oc ON oc.id = us.client_id
		WHERE oc.project_id = $1 AND us.account_id = $2
	`

	var total int64
	err := r.database.QueryRowContext(ctx, query, projectID, accountID).Scan(&total)
	if err != nil {
		return 0, stackerror.NewRepositoryError("CountSessionsByAccountIDRepository", err)
	}

	return total, nil
}

func NewCountSessionsByAccountIDRepository(database types.Database) contract.CountSessionsByAccountIDRepository {
	return &countSessionsByAccountIDRepository{database: database}
}
