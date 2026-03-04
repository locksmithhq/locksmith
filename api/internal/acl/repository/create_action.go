package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type createActionRepository struct {
	database types.Database
}

// Execute implements contract.CreateActionRepository.
func (r *createActionRepository) Execute(ctx context.Context, action domain.Action) error {
	query := "INSERT INTO actions (title) VALUES ($1)"
	_, err := r.database.ExecContext(ctx, query, action.Title)
	return err
}

func NewCreateActionRepository(database types.Database) contract.CreateActionRepository {
	return &createActionRepository{database: database}
}
