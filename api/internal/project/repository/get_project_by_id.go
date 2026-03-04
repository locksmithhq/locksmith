package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/project/contract"
	"github.com/booscaaa/locksmith/api/internal/project/domain"
)

type getProjectByIDRepository struct {
	database types.Database
}

func (r *getProjectByIDRepository) Execute(ctx context.Context, id string) (domain.Project, error) {
	var project domain.Project

	query := `SELECT id, name, description, domain FROM projects WHERE id = $1`

	err := r.database.GetContext(ctx, &project, query, id)
	if err != nil {
		return domain.Project{}, stackerror.NewRepositoryError("GetProjectByIDRepository", err)
	}

	return project, nil
}

func NewGetProjectByIDRepository(database types.Database) contract.GetProjectByIDRepository {
	return &getProjectByIDRepository{database: database}
}
