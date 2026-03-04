package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type fetchActionsRepository struct {
	database types.Database
}

// Execute implements contract.FetchRolesRepository.
func (r *fetchActionsRepository) Execute(ctx context.Context) ([]domain.Action, error) {
	var actions []domain.Action = []domain.Action{}

	if err := r.database.SelectContext(ctx, &actions, "SELECT id, title, created_at, updated_at FROM actions"); err != nil {
		return nil, err
	}

	return actions, nil
}

func NewFetchActionsRepository(database types.Database) contract.FetchActionsRepository {
	return &fetchActionsRepository{database: database}
}
