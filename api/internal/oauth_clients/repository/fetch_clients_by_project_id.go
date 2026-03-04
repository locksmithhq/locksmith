package repository

import (
	"context"

	"github.com/booscaaa/go-paginate/v3/paginate"
	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/domain"
)

type fetchClientsByProjectIDRepository struct {
	database types.Database
}

// Execute implements contract.FetchClientsByProjectIDRepository.
func (r *fetchClientsByProjectIDRepository) Execute(ctx context.Context, projectID string, params paginate.PaginationParams) ([]domain.Client, error) {
	var clients []domain.Client

	query, args, err := paginate.NewBuilder().
		Model(&domain.Client{}).
		Table("oauth_clients").
		Select("id, project_id, client_id, client_secret, redirect_uris, grant_types, name, created_at, updated_at, custom_domain").
		Where("project_id = ?", projectID).
		FromStruct(params).
		BuildSQL()

	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchClientsByProjectIDRepository", err)
	}

	err = r.database.SelectContext(ctx, &clients, query, args...)
	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchClientsByProjectIDRepository", err)
	}

	return clients, nil
}

func NewFetchClientsByProjectIDRepository(database types.Database) contract.FetchClientsByProjectIDRepository {
	return &fetchClientsByProjectIDRepository{database: database}
}
