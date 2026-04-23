package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/crypto"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/domain"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/types/input"
)

type createClientRepository struct {
	database types.Database
}

// Execute implements contract.CreateClientRepository.
func (r *createClientRepository) Execute(ctx context.Context, input input.Client) (domain.Client, error) {
	var client domain.Client

	query := `INSERT INTO oauth_clients (
		project_id, name, client_id, client_secret, redirect_uris, grant_types, custom_domain, require_pkce
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, project_id, client_id, client_secret, redirect_uris, grant_types, name, created_at, updated_at, custom_domain, require_pkce`

	encryptedSecret, err := crypto.Encrypt(input.ClientSecret)
	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("CreateClientRepository", err)
	}

	err = r.database.QueryRowxContext(ctx,
		query,
		input.ProjectID,
		input.Name,
		input.ClientID,
		encryptedSecret,
		input.RedirectURIs,
		input.GrantTypes,
		input.CustomDomain,
		input.RequirePKCE,
	).StructScan(&client)

	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("CreateClientRepository", err)
	}

	client.ClientSecret, err = crypto.Decrypt(client.ClientSecret)
	if err != nil {
		return domain.Client{}, stackerror.NewRepositoryError("CreateClientRepository", err)
	}
	return client, nil
}

func NewCreateClientRepository(database types.Database) contract.CreateClientRepository {
	return &createClientRepository{database: database}
}
