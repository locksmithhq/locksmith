package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/domain"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
)

type getAccountByProjectIDAndIDRepository struct {
	database types.Database
}

// Execute implements contract.GetAccountByProjectIDAndIDRepository.
func (r *getAccountByProjectIDAndIDRepository) Execute(ctx context.Context, projectID string, id string) (domain.Account, error) {
	var account domain.Account

	query := `SELECT 
		id, 
		project_id, 
		name, 
		email, 
		username, 
		password, 
		created_at, 
		updated_at, 
		deleted_at,
		role_name,
		must_change_password
	FROM accounts WHERE id = $1 AND project_id = $2 AND deleted_at IS NULL`
	err := r.database.GetContext(ctx, &account, query, id, projectID)
	if err != nil {
		return domain.Account{}, stackerror.NewRepositoryError("GetAccountByEmailAndProjectIDRepository", err)
	}
	return account, nil
}

func NewGetAccountByProjectIDAndIDRepository(database types.Database) contract.GetAccountByProjectIDAndIDRepository {
	return &getAccountByProjectIDAndIDRepository{database: database}
}
