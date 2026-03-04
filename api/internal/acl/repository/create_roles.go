package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/acl/contract"
	"github.com/booscaaa/locksmith/api/internal/acl/domain"
)

type createRoleRepository struct {
	database types.Database
}

// Execute implements contract.CreateRoleRepository.
func (r *createRoleRepository) Execute(ctx context.Context, role domain.Role) error {
	query := "INSERT INTO roles (title) VALUES ($1)"
	_, err := r.database.ExecContext(ctx, query, role.Title)
	return err
}

func NewCreateRoleRepository(database types.Database) contract.CreateRoleRepository {
	return &createRoleRepository{database: database}
}
