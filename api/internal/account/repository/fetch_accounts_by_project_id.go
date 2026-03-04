package repository

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/account/domain"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type fetchAccountsByProjectIDRepository struct {
	database types.Database
}

// Execute implements contract.FetchAccountsByProjectIDRepository.
func (r *fetchAccountsByProjectIDRepository) Execute(ctx context.Context, projectID string, params paginate.PaginationParams) ([]domain.Account, error) {
	var accounts []domain.Account

	query, args, err := paginate.NewBuilder().
		Model(&domain.Account{}).
		Table("accounts").
		Select("id", "name", "email", "username", "project_id", "created_at", "updated_at", "role_name", "must_change_password").
		Where("project_id = ?", projectID).
		FromStruct(params).
		BuildSQL()

	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchAccountsByProjectIDRepository", err)
	}

	err = r.database.SelectContext(ctx, &accounts, query, args...)
	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchAccountsByProjectIDRepository", err)
	}

	return accounts, nil
}

func NewFetchAccountsByProjectIDRepository(database types.Database) contract.FetchAccountsByProjectIDRepository {
	return &fetchAccountsByProjectIDRepository{database: database}
}
