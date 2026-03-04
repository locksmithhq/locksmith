package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getClientByClientIDRepository struct {
	database types.Database
}

// Execute implements contract.GetClientByClientIDRepository.
func (r *getClientByClientIDRepository) Execute(ctx context.Context, clientID string) (domain.Client, error) {
	var client domain.Client

	query := `SELECT 
		oauth_clients.id, 
		oauth_clients.project_id, 
		oauth_clients.client_id, 
		oauth_clients.client_secret, 
		oauth_clients.redirect_uris, 
		oauth_clients.grant_types, 
		oauth_clients.name, 
		oauth_clients.created_at, 
		oauth_clients.updated_at, 
		projects.domain
	FROM oauth_clients 
	INNER JOIN projects ON oauth_clients.project_id = projects.id
	WHERE oauth_clients.client_id = $1`

	err := r.database.QueryRowxContext(ctx, query, clientID).StructScan(&client)
	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("GetClientByClientIDRepository", err)
	}
	return client, nil
}

func NewGetClientByClientIDRepository(database types.Database) contract.GetClientByClientIDRepository {
	return &getClientByClientIDRepository{database: database}
}
