package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/google/uuid"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type createAccountRepository struct {
	database types.Database
}

func (r *createAccountRepository) Execute(ctx context.Context, account domain.Account) (domain.Account, error) {
	var result domain.Account

	id := uuid.New().String()

	query := `INSERT INTO accounts (id, project_id, name, email, username, password, role_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, '', $6, NOW(), NOW())
		RETURNING id, project_id, name, email, username, password, role_name, created_at, updated_at, deleted_at, must_change_password`

	err := r.database.QueryRowxContext(ctx, query,
		id,
		account.ProjectID,
		account.Name,
		account.Email,
		account.Username,
		account.RoleName,
	).StructScan(&result)
	if err != nil {
		return domain.Account{}, stackerror.NewRepositoryError("CreateAccountRepository", err)
	}
	return result, nil
}

func NewCreateAccountRepository(database types.Database) contract.CreateAccountRepository {
	return &createAccountRepository{database: database}
}
