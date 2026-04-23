package repository

import (
	"context"
	"time"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type revokeUserSessionRepository struct {
	database types.Database
}

func (r *revokeUserSessionRepository) Execute(ctx context.Context, sessionID string) error {
	now := time.Now().Format(time.RFC3339)
	query := `UPDATE user_sessions SET revoked = true, revoked_at = $1, revoked_reason = 'user_logout' WHERE id = $2`
	_, err := r.database.ExecContext(ctx, query, now, sessionID)
	if err != nil {
		return stackerror.NewRepositoryError("RevokeUserSessionRepository", err)
	}
	return nil
}

func NewRevokeUserSessionRepository(database types.Database) contract.RevokeUserSessionRepository {
	return &revokeUserSessionRepository{database: database}
}
