package repository

import (
	"context"
	"database/sql"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/hasher"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
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
		FROM accounts WHERE email = $1 AND project_id = $2
	`
	err := r.database.GetContext(ctx, &account, query, email, projectID)
	if err != nil {
		return domain.Account{}, stackerror.NewRepositoryError("GetAccountByEmailPasswordAndProjectIDRepository", err)
	}

	ok, err := hasher.Verify(password, account.Password)
	if err != nil || !ok {
		return domain.Account{}, stackerror.NewRepositoryError(
			"GetAccountByEmailPasswordAndProjectIDRepository",
			sql.ErrNoRows,
		)
	}

	return account, nil
}

func NewGetAccountByEmailPasswordAndProjectIDRepository(database types.Database) contract.GetAccountByEmailPasswordAndProjectIDRepository {
	return &getAccountByEmailPasswordAndProjectIDRepository{database: database}
}
