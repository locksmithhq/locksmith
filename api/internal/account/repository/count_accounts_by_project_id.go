package repository

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/account/contract"
	"github.com/booscaaa/locksmith/api/internal/account/domain"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
)

type countAccountsByProjectIDRepository struct {
	database types.Database
}

// Execute implements contract.CountAccountsByProjectIDRepository.
func (r *countAccountsByProjectIDRepository) Execute(ctx context.Context, projectID string, params paginate.PaginationParams) (int64, error) {
	query, args, err := paginate.NewBuilder().
		Model(&domain.Account{}).
		Table("accounts").
		Where("project_id = ?", projectID).
		FromStruct(params).
		BuildCountSQL()

	if err != nil {
		return 0, stackerror.NewRepositoryError("CountAccountsByProjectIDRepository", err)
	}

	var total int64
	err = r.database.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return 0, stackerror.NewRepositoryError("CountAccountsByProjectIDRepository", err)
	}

	return total, nil
}

func NewCountAccountsByProjectIDRepository(database types.Database) contract.CountAccountsByProjectIDRepository {
	return &countAccountsByProjectIDRepository{database: database}
}
