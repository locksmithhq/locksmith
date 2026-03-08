package repository

import (
	"context"
	"time"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type updateUserSessionActivityRepository struct {
	database types.Database
}

func (r *updateUserSessionActivityRepository) Execute(ctx context.Context, sessionID string) error {
	now := time.Now()
	query := `UPDATE user_sessions SET last_activity = $1, expires_at = $2 WHERE id = $3`
	_, err := r.database.ExecContext(ctx, query, now.Format(time.RFC3339), now.Add(24*time.Hour).Format(time.RFC3339), sessionID)
	if err != nil {
		return stackerror.NewRepositoryError("UpdateUserSessionActivityRepository", err)
	}
	return nil
}

func NewUpdateUserSessionActivityRepository(database types.Database) contract.UpdateUserSessionActivityRepository {
	return &updateUserSessionActivityRepository{database: database}
}
