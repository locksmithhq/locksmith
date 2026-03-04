package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/domain"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
)

type updateAccountRepository struct {
	database types.Database
}

// Execute implements contract.UpdateAccountRepository.
func (r *updateAccountRepository) Execute(ctx context.Context, entity domain.Account) (domain.Account, error) {
	var account domain.Account

	query := `UPDATE accounts SET 
		name = $1, 
		email = $2, 
		username = $3, 
		role_name = $4,
		password = CASE WHEN $5 = '' THEN password ELSE crypt($5, gen_salt('bf', 8)) END,
		must_change_password = $6,
		updated_at = NOW()
	WHERE id = $7 AND project_id = $8
	RETURNING id, project_id, name, email, username, role_name, created_at, updated_at, must_change_password`

	err := r.database.QueryRowxContext(ctx,
		query,
		entity.Name,
		entity.Email,
		entity.Username,
		entity.RoleName,
		entity.Password,
		entity.MustChangePassword,
		entity.ID,
		entity.ProjectID,
	).StructScan(&account)

	if err != nil {
		return domain.Account{}, stackerror.NewRepositoryError("UpdateAccountRepository", err)
	}
	return account, nil
}

func NewUpdateAccountRepository(database types.Database) contract.UpdateAccountRepository {
	return &updateAccountRepository{database: database}
}
