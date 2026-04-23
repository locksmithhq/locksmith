package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type createSocialStateRepository struct {
	database types.Database
}

func (r *createSocialStateRepository) Execute(ctx context.Context, state domain.SocialState) error {
	query := `INSERT INTO oauth_social_states (nonce, client_id, redirect_uri, state, code_challenge, code_challenge_method, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.database.ExecContext(ctx, query,
		state.Nonce,
		state.ClientID,
		state.RedirectURI,
		state.State,
		state.CodeChallenge,
		state.CodeChallengeMethod,
		state.ExpiresAt,
	)
	if err != nil {
		return stackerror.NewRepositoryError("CreateSocialStateRepository", err)
	}
	return nil
}

func NewCreateSocialStateRepository(database types.Database) contract.CreateSocialStateRepository {
	return &createSocialStateRepository{database: database}
}
