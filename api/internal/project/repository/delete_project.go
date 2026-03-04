package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
)

type deleteProjectRepository struct {
	database types.Database
}

func (r *deleteProjectRepository) Execute(ctx context.Context, id string) error {
	query := "UPDATE projects SET deleted_at = NOW() WHERE id = $1"

	_, err := r.database.ExecContext(ctx, query, id)
	if err != nil {
		return stackerror.NewRepositoryError("DeleteProjectRepository", err)
	}

	return nil
}

func NewDeleteProjectRepository(database types.Database) contract.DeleteProjectRepository {
	return &deleteProjectRepository{database: database}
}
