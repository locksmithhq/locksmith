package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth2/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth2/domain"
)

type updateRefreshTokenRepository struct {
	database types.Database
}

// Execute implements contract.UpdateRefreshTokenRepository.
func (r *updateRefreshTokenRepository) Execute(ctx context.Context, entity domain.RefreshToken) error {
	query := `UPDATE refresh_tokens SET 
		revoked = $1, 
		revoked_at = $2, 
		revoked_reason = $3, 
		last_used_at = $4,
		rotation_count = $5
		WHERE id = $6`

	_, err := r.database.ExecContext(ctx,
		query,
		entity.Revoked,
		entity.RevokedAt,
		entity.RevokedReason,
		entity.LastUsedAt,
		entity.RotationCount,
		entity.ID,
	)
	if err != nil {
		return stackerror.NewRepositoryError("UpdateRefreshTokenRepository", err)
	}
	return nil
}

func NewUpdateRefreshTokenRepository(database types.Database) contract.UpdateRefreshTokenRepository {
	return &updateRefreshTokenRepository{database: database}
}
