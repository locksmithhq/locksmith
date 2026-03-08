package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getAuthCodeByCodeRepository struct {
	database types.Database
}

// Execute implements contract.GetAuthCodeByCodeRepository.
func (r *getAuthCodeByCodeRepository) Execute(ctx context.Context, code string) (domain.AuthCode, error) {
	var authCode domain.AuthCode

	query := `SELECT id, code, client_id, account_id, redirect_uri, code_challenge, code_challenge_method, expires_at, used, created_at 
				FROM oauth_authorization_codes 
				WHERE code = $1 LIMIT 1`

	err := r.database.GetContext(ctx, &authCode, query, code)
	if err != nil {
		return domain.AuthCode{}, stackerror.NewRepositoryError("GetAuthCodeByCodeRepository", err)
	}
	return authCode, nil
}

func NewGetAuthCodeByCodeRepository(database types.Database) contract.GetAuthCodeByCodeRepository {
	return &getAuthCodeByCodeRepository{database: database}
}
