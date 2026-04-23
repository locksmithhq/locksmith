package repository

import (
	"context"
	"time"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type revokeRefreshTokensBySessionRepository struct {
	database types.Database
}

func (r *revokeRefreshTokensBySessionRepository) Execute(ctx context.Context, sessionID string) error {
	now := time.Now().Format(time.RFC3339)
	query := `UPDATE refresh_tokens
		SET revoked = true, revoked_at = $1, revoked_reason = 'token_reuse_detected'
		WHERE session_id = $2 AND revoked = false`
	_, err := r.database.ExecContext(ctx, query, now, sessionID)
	if err != nil {
		return stackerror.NewRepositoryError("RevokeRefreshTokensBySessionRepository", err)
	}
	return nil
}

func NewRevokeRefreshTokensBySessionRepository(database types.Database) contract.RevokeRefreshTokensBySessionRepository {
	return &revokeRefreshTokensBySessionRepository{database: database}
}
