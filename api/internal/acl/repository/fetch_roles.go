package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type fetchRolesRepository struct {
	database types.Database
}

// Execute implements contract.FetchRolesRepository.
func (r *fetchRolesRepository) Execute(ctx context.Context) ([]domain.Role, error) {
	var roles []domain.Role = []domain.Role{}

	if err := r.database.SelectContext(ctx, &roles, "SELECT id, title, created_at, updated_at FROM roles"); err != nil {
		return nil, err
	}

	return roles, nil
}

func NewFetchRolesRepository(database types.Database) contract.FetchRolesRepository {
	return &fetchRolesRepository{database: database}
}
