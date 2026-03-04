package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/locksmithhq/locksmith/api/internal/project/domain"
)

type getProjectByNameRepository struct {
	database types.Database
}

func (r *getProjectByNameRepository) Execute(ctx context.Context, name string) (domain.Project, error) {
	var project domain.Project

	query := `SELECT id, name, description, domain FROM projects WHERE name = $1`

	err := r.database.GetContext(ctx, &project, query, name)
	if err != nil {
		return domain.Project{}, stackerror.NewRepositoryError("GetProjectByNameRepository", err)
	}

	return project, nil
}

func NewGetProjectByNameRepository(database types.Database) contract.GetProjectByNameRepository {
	return &getProjectByNameRepository{database: database}
}
