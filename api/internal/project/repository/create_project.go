package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/project/contract"
	"github.com/booscaaa/locksmith/api/internal/project/domain"
)

type createProjectRepository struct {
	database types.Database
}

// Execute implements contract.CreateProjectRepository.
func (repository *createProjectRepository) Execute(ctx context.Context, entity domain.Project) (domain.Project, error) {
	var project domain.Project

	query := "INSERT INTO projects (name, description, domain) VALUES ($1, $2, $3) RETURNING id, name, description, domain"

	err := repository.database.QueryRowxContext(
		ctx,
		query,
		entity.Name,
		entity.Description,
		entity.Domain,
	).StructScan(&project)
	if err != nil {
		return domain.Project{}, stackerror.NewRepositoryError("CreateProjectRepository", err)
	}

	return project, nil
}

func NewCreateProjectRepository(database types.Database) contract.CreateProjectRepository {
	return &createProjectRepository{database: database}
}
