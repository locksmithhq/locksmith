package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/core/hasher"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
)

type updateAccountPasswordRepository struct {
	database types.Database
}

// Execute implements contract.UpdateAccountPasswordRepository.
func (r *updateAccountPasswordRepository) Execute(ctx context.Context, id string, password string) error {
	hashedPassword, err := hasher.Hash(password)
	if err != nil {
		return stackerror.NewRepositoryError("UpdateAccountPasswordRepository", err)
	}

	query := `UPDATE accounts SET
		password = $1,
		must_change_password = false,
		updated_at = NOW()
	WHERE id = $2`

	_, err = r.database.ExecContext(ctx, query, hashedPassword, id)
	if err != nil {
		return stackerror.NewRepositoryError("UpdateAccountPasswordRepository", err)
	}

	return nil
}

func NewUpdateAccountPasswordRepository(database types.Database) contract.UpdateAccountPasswordRepository {
	return &updateAccountPasswordRepository{database: database}
}
