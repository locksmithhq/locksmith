package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type createRefreshTokenRepository struct {
	database types.Database
}

// Execute implements contract.CreateRefreshTokenRepository.
func (r *createRefreshTokenRepository) Execute(ctx context.Context, entity domain.RefreshToken) (domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken

	query := `INSERT INTO refresh_tokens (
		token_hash, session_id, account_id, client_id, expires_at, rotation_count, parent_token_id
	) VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id, token_hash, session_id, account_id, client_id, expires_at, rotation_count, parent_token_id, created_at`

	err := r.database.QueryRowxContext(ctx,
		query,
		entity.TokenHash,
		entity.SessionID,
		entity.AccountID,
		entity.ClientID,
		entity.ExpiresAt,
		entity.RotationCount,
		entity.ParentTokenID,
	).StructScan(&refreshToken)

	if err != nil {
		return domain.RefreshToken{}, stackerror.NewRepositoryError("CreateRefreshTokenRepository", err)
	}
	return refreshToken, nil
}

func NewCreateRefreshTokenRepository(database types.Database) contract.CreateRefreshTokenRepository {
	return &createRefreshTokenRepository{database: database}
}
