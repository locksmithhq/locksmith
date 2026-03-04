package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getRefreshTokenByHashRepository struct {
	database types.Database
}

// Execute implements contract.GetRefreshTokenByHashRepository.
func (r *getRefreshTokenByHashRepository) Execute(ctx context.Context, hash string) (domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	query := `SELECT 
			id,
			token_hash,
			session_id,
			account_id,
			client_id,
			rotation_count,
			parent_token_id,
			expires_at,
			revoked,
			revoked_at,
			revoked_reason,
			last_used_at,
			created_at
		FROM refresh_tokens WHERE token_hash = $1`
	err := r.database.QueryRowxContext(ctx,
		query,
		hash,
	).StructScan(&refreshToken)

	if err != nil {
		return domain.RefreshToken{}, stackerror.NewRepositoryError("GetRefreshTokenByHashRepository", err)
	}
	return refreshToken, nil
}

func NewGetRefreshTokenByHashRepository(database types.Database) contract.GetRefreshTokenByHashRepository {
	return &getRefreshTokenByHashRepository{database: database}
}
