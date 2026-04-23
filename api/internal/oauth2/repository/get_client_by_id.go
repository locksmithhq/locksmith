package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/crypto"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getClientByIDRepository struct {
	database types.Database
}

// Execute implements contract.GetClientByIDRepository.
func (r *getClientByIDRepository) Execute(ctx context.Context, id string) (domain.Client, error) {
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
		oauth_clients.require_pkce,
		projects.domain
	FROM oauth_clients
	INNER JOIN projects ON oauth_clients.project_id = projects.id
	WHERE oauth_clients.id = $1`

	err := r.database.QueryRowxContext(ctx, query, id).StructScan(&client)
	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("GetClientByIDRepository", err)
	}

	client.ClientSecret, err = crypto.Decrypt(client.ClientSecret)
	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("GetClientByIDRepository", err)
	}
	return client, nil
}

func NewGetClientByIDRepository(database types.Database) contract.GetClientByIDRepository {
	return &getClientByIDRepository{database: database}
}
