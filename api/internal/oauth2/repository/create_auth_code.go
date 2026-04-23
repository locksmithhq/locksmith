package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type createAuthCodeRepository struct {
	database types.Database
}

// Execute implements contract.CreateAuthCodeRepository.
func (r *createAuthCodeRepository) Execute(ctx context.Context, entity domain.AuthCode) (domain.AuthCode, error) {
	var authCode domain.AuthCode

	query := `INSERT INTO oauth_authorization_codes (
		code, client_id, account_id, redirect_uri, code_challenge, code_challenge_method, expires_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id, code, client_id, account_id, redirect_uri, code_challenge, code_challenge_method, expires_at, used, created_at`

	h := sha256.Sum256([]byte(entity.Code))
	codeHash := hex.EncodeToString(h[:])

	err := r.database.QueryRowxContext(ctx,
		query,
		codeHash,
		entity.ClientID,
		entity.AccountID,
		entity.RedirectURI,
		entity.CodeChallenge,
		entity.CodeChallengeMethod,
		entity.ExpiresAt,
	).StructScan(&authCode)

	if err != nil {
		return domain.AuthCode{}, stackerror.NewRepositoryError("CreateAuthCodeRepository", err)
	}

	// Return the raw code so the calling usecase can embed it in the redirect URL.
	// The hash is what lives in the database.
	authCode.Code = entity.Code
	return authCode, nil
}

func NewCreateAuthCodeRepository(database types.Database) contract.CreateAuthCodeRepository {
	return &createAuthCodeRepository{database: database}
}
