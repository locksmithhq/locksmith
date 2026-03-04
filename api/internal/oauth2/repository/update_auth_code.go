package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type updateAuthCodeRepository struct {
	database types.Database
}

// Execute implements contract.UpdateAuthCodeRepository.
func (r *updateAuthCodeRepository) Execute(ctx context.Context, entity domain.AuthCode) error {
	query := `UPDATE oauth_authorization_codes 
	          SET used = $1 
	          WHERE id = $2`

	_, err := r.database.ExecContext(ctx, query, entity.Used, entity.ID)
	if err != nil {
		return stackerror.NewRepositoryError("UpdateAuthCodeRepository", err)
	}
	return nil
}

func NewUpdateAuthCodeRepository(database types.Database) contract.UpdateAuthCodeRepository {
	return &updateAuthCodeRepository{database: database}
}
