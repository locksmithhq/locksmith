package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type deleteSocialStateRepository struct {
	database types.Database
}

func (r *deleteSocialStateRepository) Execute(ctx context.Context, nonce string) error {
	_, err := r.database.ExecContext(ctx, `DELETE FROM oauth_social_states WHERE nonce = $1`, nonce)
	if err != nil {
		return stackerror.NewRepositoryError("DeleteSocialStateRepository", err)
	}
	return nil
}

func NewDeleteSocialStateRepository(database types.Database) contract.DeleteSocialStateRepository {
	return &deleteSocialStateRepository{database: database}
}
