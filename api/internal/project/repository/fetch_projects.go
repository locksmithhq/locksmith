package repository

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/project/contract"
	"github.com/booscaaa/locksmith/api/internal/project/domain"
)

type fetchProjectsRepository struct {
	database types.Database
}

// Execute implements contract.FetchProjectsRepository.
func (r *fetchProjectsRepository) Execute(ctx context.Context, params paginate.PaginationParams) ([]domain.Project, error) {
	var projects []domain.Project
	query, args, err := paginate.NewBuilder().
		Model(&domain.Project{}).
		Table("projects").
		Select("id, name, description, domain").
		Where("deleted_at IS NULL").
		FromStruct(params).
		BuildSQL()

	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchProjectsRepository", err)
	}

	err = r.database.SelectContext(ctx, &projects, query, args...)
	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchProjectsRepository", err)
	}

	return projects, nil
}

func NewFetchProjectsRepository(database types.Database) contract.FetchProjectsRepository {
	return &fetchProjectsRepository{database: database}
}
