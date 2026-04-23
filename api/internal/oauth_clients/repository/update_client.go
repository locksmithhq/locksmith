package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/crypto"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/domain"
)

type updateClientRepository struct {
	database types.Database
}

// Execute implements contract.UpdateClientRepository.
func (r *updateClientRepository) Execute(ctx context.Context, projectID string, id string, in domain.Client) (domain.Client, error) {
	var client domain.Client

	query := `UPDATE oauth_clients
		SET name = $1, client_secret = $2, redirect_uris = $3, grant_types = $4, custom_domain = $5, require_pkce = $6
		WHERE project_id = $7 AND id = $8
		RETURNING id, project_id, client_id, client_secret, redirect_uris, grant_types, name, created_at, updated_at, custom_domain, require_pkce`

	encryptedSecret, err := crypto.Encrypt(in.ClientSecret)
	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("UpdateClientRepository", err)
	}

	err = r.database.QueryRowxContext(ctx,
		query,
		in.Name,
		encryptedSecret,
		in.RedirectURIs,
		in.GrantTypes,
		in.CustomDomain,
		in.RequirePKCE,
		projectID,
		id,
	).StructScan(&client)

	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("UpdateClientRepository", err)
	}

	client.ClientSecret, err = crypto.Decrypt(client.ClientSecret)
	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("UpdateClientRepository", err)
	}
	return client, nil
}

func NewUpdateClientRepository(database types.Database) contract.UpdateClientRepository {
	return &updateClientRepository{database: database}
}
