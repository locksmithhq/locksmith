package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients/domain"
)

type getClientByIDAndProjectIDRepository struct {
	database types.Database
}

// Execute implements contract.GetClientByIDAndProjectIDRepository.
func (r *getClientByIDAndProjectIDRepository) Execute(ctx context.Context, id string, projectID string) (domain.Client, error) {
	var client domain.Client

	query := `
		SELECT 
			id, 
			project_id, 
			client_id, 
			client_secret, 
			redirect_uris, 
			grant_types, 
			name, 
			created_at, 
			updated_at,
			custom_domain
		FROM oauth_clients WHERE id = $1 AND project_id = $2
	`
	err := r.database.GetContext(ctx, &client, query, id, projectID)
	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("GetClientByIDAndProjectIDRepository", err)
	}

	return client, nil
}

func NewGetClientByIDAndProjectIDRepository(database types.Database) contract.GetClientByIDAndProjectIDRepository {
	return &getClientByIDAndProjectIDRepository{
		database: database,
	}
}
