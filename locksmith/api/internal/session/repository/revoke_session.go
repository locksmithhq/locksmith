package repository

import (
	"context"
	"database/sql"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
)

type revokeSessionRepository struct {
	database types.Database
}

func (r *revokeSessionRepository) Execute(ctx context.Context, projectID, sessionID string) error {
	query := `
		UPDATE user_sessions us SET
			revoked = true,
			revoked_at = NOW(),
			revoked_reason = 'manual_revoke'
		FROM oauth_clients oc
		WHERE us.client_id = oc.id
			AND oc.project_id = $1
			AND us.id = $2
			AND us.revoked = false
	`
	result, err := r.database.ExecContext(ctx, query, projectID, sessionID)
	if err != nil {
		return stackerror.NewRepositoryError("RevokeSessionRepository", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return stackerror.NewRepositoryError("RevokeSessionRepository", err)
	}
	if rows == 0 {
		return stackerror.NewRepositoryError("RevokeSessionRepository", sql.ErrNoRows)
	}

	// Revoke all refresh tokens linked to this session
	revokeTokensQuery := `
		UPDATE refresh_tokens SET
			revoked = true,
			revoked_at = NOW(),
			revoked_reason = 'session_revoked'
		WHERE session_id = $1 AND revoked = false
	`
	_, err = r.database.ExecContext(ctx, revokeTokensQuery, sessionID)
	if err != nil {
		return stackerror.NewRepositoryError("RevokeSessionRepository", err)
	}

	return nil
}

func NewRevokeSessionRepository(database types.Database) contract.RevokeSessionRepository {
	return &revokeSessionRepository{database: database}
}
