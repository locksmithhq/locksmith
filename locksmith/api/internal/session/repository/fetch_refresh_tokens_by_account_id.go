package repository

import (
	"context"
	"fmt"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/locksmithhq/locksmith/api/internal/session/domain"
)

type fetchRefreshTokensByAccountIDRepository struct {
	database types.Database
}

func (r *fetchRefreshTokensByAccountIDRepository) Execute(ctx context.Context, projectID, accountID, sessionID string, page, limit int) ([]domain.RefreshToken, error) {
	var tokens []domain.RefreshToken
	offset := (page - 1) * limit

	query := `
		SELECT
			rt.id,
			COALESCE(rt.session_id::text, '') AS session_id,
			rt.account_id,
			rt.client_id,
			rt.rotation_count,
			COALESCE(rt.parent_token_id::text, '') AS parent_token_id,
			to_char(rt.expires_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS expires_at,
			rt.revoked,
			COALESCE(to_char(rt.revoked_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"'), '') AS revoked_at,
			COALESCE(rt.revoked_reason, '') AS revoked_reason,
			COALESCE(to_char(rt.last_used_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"'), '') AS last_used_at,
			to_char(rt.created_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS created_at,
			oc.name AS client_name
		FROM refresh_tokens rt
		JOIN oauth_clients oc ON oc.id = rt.client_id
		WHERE oc.project_id = $1 AND rt.account_id = $2
	`

	args := []interface{}{projectID, accountID}
	argIdx := 3

	if sessionID != "" {
		query += fmt.Sprintf(" AND rt.session_id = $%d", argIdx)
		args = append(args, sessionID)
		argIdx++
	}

	query += fmt.Sprintf(" ORDER BY rt.created_at DESC LIMIT $%d OFFSET $%d", argIdx, argIdx+1)
	args = append(args, limit, offset)

	err := r.database.SelectContext(ctx, &tokens, query, args...)
	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchRefreshTokensByAccountIDRepository", err)
	}

	return tokens, nil
}

func NewFetchRefreshTokensByAccountIDRepository(database types.Database) contract.FetchRefreshTokensByAccountIDRepository {
	return &fetchRefreshTokensByAccountIDRepository{database: database}
}
