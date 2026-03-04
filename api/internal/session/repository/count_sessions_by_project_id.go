package repository

import (
	"context"
	"fmt"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/session/contract"
)

type countSessionsByProjectIDRepository struct {
	database types.Database
}

func (r *countSessionsByProjectIDRepository) Execute(ctx context.Context, projectID string, search string) (int64, error) {
	query := `
		SELECT COUNT(*)
		FROM user_sessions us
		JOIN accounts a ON a.id = us.account_id
		JOIN oauth_clients oc ON oc.id = us.client_id
		WHERE oc.project_id = $1
	`

	args := []interface{}{projectID}

	if search != "" {
		query += fmt.Sprintf(
			` AND (a.name ILIKE $%d OR a.email ILIKE $%d OR COALESCE(us.ip_address::text, '') ILIKE $%d OR COALESCE(us.browser, '') ILIKE $%d)`,
			2, 2, 2, 2,
		)
		args = append(args, "%"+search+"%")
	}

	var total int64
	err := r.database.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return 0, stackerror.NewRepositoryError("CountSessionsByProjectIDRepository", err)
	}

	return total, nil
}

func NewCountSessionsByProjectIDRepository(database types.Database) contract.CountSessionsByProjectIDRepository {
	return &countSessionsByProjectIDRepository{database: database}
}
