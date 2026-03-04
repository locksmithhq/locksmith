package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth2/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth2/domain"
)

type getAccountByEmailPasswordAndProjectIDRepository struct {
	database types.Database
}

// Execute implements contract.GetAccountByEmailPasswordAndProjectIDRepository.
func (r *getAccountByEmailPasswordAndProjectIDRepository) Execute(ctx context.Context, email string, password string, projectID string) (domain.Account, error) {
	var account domain.Account
	query := `
		SELECT 
			id, 
			project_id, 
			name, 
			email, 
			username, 
			password, 
			created_at, 
			updated_at, 
			deleted_at,
			must_change_password
		FROM accounts WHERE email = $1 AND password = crypt($2, password) AND project_id = $3
	`
	err := r.database.GetContext(ctx, &account, query, email, password, projectID)
	if err != nil {
		return domain.Account{}, stackerror.NewRepositoryError("GetAccountByEmailPasswordAndProjectIDRepository", err)
	}
	return account, nil
}

func NewGetAccountByEmailPasswordAndProjectIDRepository(database types.Database) contract.GetAccountByEmailPasswordAndProjectIDRepository {
	return &getAccountByEmailPasswordAndProjectIDRepository{database: database}
}
