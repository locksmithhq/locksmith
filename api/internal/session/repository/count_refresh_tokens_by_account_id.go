package repository

import (
	"context"
	"fmt"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type countRefreshTokensByAccountIDRepository struct {
	database types.Database
}

func (r *countRefreshTokensByAccountIDRepository) Execute(ctx context.Context, projectID, accountID, sessionID string) (int64, error) {
	var count int64

	query := `
		SELECT COUNT(rt.id)
		FROM refresh_tokens rt
		JOIN oauth_clients oc ON oc.id = rt.client_id
		WHERE oc.project_id = $1 AND rt.account_id = $2
	`

	args := []interface{}{projectID, accountID}
	argIdx := 3

	if sessionID != "" {
		query += fmt.Sprintf(" AND rt.session_id = $%d", argIdx)
		args = append(args, sessionID)
	}

	err := r.database.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, stackerror.NewRepositoryError("CountRefreshTokensByAccountIDRepository", err)
	}

	return count, nil
}

func NewCountRefreshTokensByAccountIDRepository(database types.Database) contract.CountRefreshTokensByAccountIDRepository {
	return &countRefreshTokensByAccountIDRepository{database: database}
}
