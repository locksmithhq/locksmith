package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getSocialStateByNonceRepository struct {
	database types.Database
}

func (r *getSocialStateByNonceRepository) Execute(ctx context.Context, nonce string) (domain.SocialState, error) {
	var result domain.SocialState

	query := `SELECT nonce, client_id, redirect_uri, state, code_challenge, code_challenge_method, expires_at
		FROM oauth_social_states WHERE nonce = $1 AND expires_at > NOW()`

	err := r.database.QueryRowxContext(ctx, query, nonce).StructScan(&result)
	if err != nil {
		return domain.SocialState{}, stackerror.NewRepositoryError("GetSocialStateByNonceRepository", err)
	}
	return result, nil
}

func NewGetSocialStateByNonceRepository(database types.Database) contract.GetSocialStateByNonceRepository {
	return &getSocialStateByNonceRepository{database: database}
}
