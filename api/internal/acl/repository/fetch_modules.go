package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type fetchModulesRepository struct {
	database types.Database
}

// Execute implements contract.FetchRolesRepository.
func (r *fetchModulesRepository) Execute(ctx context.Context) ([]domain.Module, error) {
	var modules []domain.Module = []domain.Module{}

	if err := r.database.SelectContext(ctx, &modules, "SELECT id, title, created_at, updated_at FROM modules"); err != nil {
		return nil, err
	}

	return modules, nil
}

func NewFetchModulesRepository(database types.Database) contract.FetchModulesRepository {
	return &fetchModulesRepository{database: database}
}
