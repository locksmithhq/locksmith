package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/account/contract"
	"github.com/locksmithhq/locksmith/api/internal/account/domain"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
)

type getAccountByEmailAndProjectIDRepository struct {
	database types.Database
}

// Execute implements contract.GetAccountByEmailAndProjectIDRepository.
func (r *getAccountByEmailAndProjectIDRepository) Execute(ctx context.Context, email string, projectID string) (domain.Account, error) {
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
		deleted_at 
	FROM accounts WHERE email = $1 AND project_id = $2`
	err := r.database.GetContext(ctx, &account, query, email, projectID)
	if err != nil {
		return domain.Account{}, stackerror.NewRepositoryError("GetAccountByEmailAndProjectIDRepository", err)
	}
	return account, nil
}

func NewGetAccountByEmailAndProjectIDRepository(database types.Database) contract.GetAccountByEmailAndProjectIDRepository {
	return &getAccountByEmailAndProjectIDRepository{database: database}
}
