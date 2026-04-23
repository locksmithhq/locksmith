package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getAccountByEmailAndProjectRepository struct {
	database types.Database
}

func (r *getAccountByEmailAndProjectRepository) Execute(ctx context.Context, email string, projectID string) (domain.Account, error) {
	var account domain.Account

	query := `SELECT id, project_id, name, email, username, password, created_at, updated_at, deleted_at, must_change_password
		FROM accounts WHERE email = $1 AND project_id = $2 AND deleted_at IS NULL`

	err := r.database.QueryRowxContext(ctx, query, email, projectID).StructScan(&account)
	if err != nil {
		return domain.Account{}, stackerror.NewRepositoryError("GetAccountByEmailAndProjectRepository", err)
	}
	return account, nil
}

func NewGetAccountByEmailAndProjectRepository(database types.Database) contract.GetAccountByEmailAndProjectRepository {
	return &getAccountByEmailAndProjectRepository{database: database}
}
