package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type updateAccountPasswordRepository struct {
	database types.Database
}

// Execute implements contract.UpdateAccountPasswordRepository.
func (r *updateAccountPasswordRepository) Execute(ctx context.Context, id string, password string) error {
	query := `UPDATE accounts SET 
		password = crypt($1, gen_salt('bf', 8)),
		must_change_password = false,
		updated_at = NOW()
	WHERE id = $2`

	_, err := r.database.ExecContext(ctx, query, password, id)
	if err != nil {
		return stackerror.NewRepositoryError("UpdateAccountPasswordRepository", err)
	}

	return nil
}

func NewUpdateAccountPasswordRepository(database types.Database) contract.UpdateAccountPasswordRepository {
	return &updateAccountPasswordRepository{database: database}
}
