package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/project/contract"
	"github.com/locksmithhq/locksmith/api/internal/project/domain"
)

type updateProjectRepository struct {
	database types.Database
}

// Execute implements contract.UpdateProjectRepository.
func (repository *updateProjectRepository) Execute(ctx context.Context, id string, entity domain.Project) (domain.Project, error) {
	var project domain.Project

	query := "UPDATE projects SET name = $1, description = $2 WHERE id = $3 RETURNING id, name, description, domain"

	err := repository.database.QueryRowxContext(
		ctx,
		query,
		entity.Name,
		entity.Description,
		id,
	).StructScan(&project)
	if err != nil {
		return domain.Project{}, stackerror.NewRepositoryError("UpdateProjectRepository", err)
	}

	return project, nil
}

func NewUpdateProjectRepository(database types.Database) contract.UpdateProjectRepository {
	return &updateProjectRepository{database: database}
}
