package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/account/domain"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type createAccountRepository struct {
	database types.Database
}

// Execute implements contract.CreateAccountRepository.
func (r *createAccountRepository) Execute(ctx context.Context, entity domain.Account) (domain.Account, error) {
	var account domain.Account

	query := `INSERT INTO accounts (
		project_id, name, email, username, password, role_name, must_change_password
	) VALUES ($1, $2, $3, $4, crypt($5, gen_salt('bf', 8)), $6, $7) RETURNING 
		id, project_id, name, email, username, password, created_at, updated_at, must_change_password`
	err := r.database.QueryRowxContext(ctx,
		query,
		entity.ProjectID,
		entity.Name,
		entity.Email,
		entity.Username,
		entity.Password,
		entity.RoleName,
		entity.MustChangePassword,
	).StructScan(&account)
	if err != nil {
		return domain.Account{}, stackerror.NewRepositoryError("CreateAccountRepository", err)
	}
	return account, nil
}

func NewCreateAccountRepository(database types.Database) contract.CreateAccountRepository {
	return &createAccountRepository{database: database}
}
