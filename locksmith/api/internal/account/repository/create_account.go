package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/domain"
	"github.com/locksmithhq/locksmith/api/internal/core/hasher"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
)

type createAccountRepository struct {
	database types.Database
}

// Execute implements contract.CreateAccountRepository.
func (r *createAccountRepository) Execute(ctx context.Context, entity domain.Account) (domain.Account, error) {
	var account domain.Account

	hashedPassword, err := hasher.Hash(entity.Password)
	if err != nil {
		return domain.Account{}, stackerror.NewRepositoryError("CreateAccountRepository", err)
	}

	query := `INSERT INTO accounts (
		project_id, name, email, username, password, role_name, must_change_password
	) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING
		id, project_id, name, email, username, password, created_at, updated_at, must_change_password`
	err = r.database.QueryRowxContext(ctx,
		query,
		entity.ProjectID,
		entity.Name,
		entity.Email,
		entity.Username,
		hashedPassword,
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
