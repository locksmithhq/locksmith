package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type createModuleRepository struct {
	database types.Database
}

// Execute implements contract.CreateModuleRepository.
func (r *createModuleRepository) Execute(ctx context.Context, module domain.Module) error {
	query := "INSERT INTO modules (title) VALUES ($1)"
	_, err := r.database.ExecContext(ctx, query, module.Title)
	return err
}

func NewCreateModuleRepository(database types.Database) contract.CreateModuleRepository {
	return &createModuleRepository{database: database}
}
