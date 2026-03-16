package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/contract"
)

type getClientByCustomDomainRepository struct {
	database types.Database
}

// Execute implements contract.GetClientByCustomDomainRepository.
func (r *getClientByCustomDomainRepository) Execute(ctx context.Context, domain string) (string, string, error) {
	var row struct {
		ClientID     string `db:"client_id"`
		RedirectURIs string `db:"redirect_uris"`
	}

	query := `
		SELECT client_id, redirect_uris FROM oauth_clients
		WHERE custom_domain = $1
		LIMIT 1
	`
	err := r.database.GetContext(ctx, &row, query, domain)
	if err != nil {
		return "", "", stackerror.NewRepositoryError("GetClientByCustomDomainRepository", err)
	}

	return row.ClientID, row.RedirectURIs, nil
}

func NewGetClientByCustomDomainRepository(database types.Database) contract.GetClientByCustomDomainRepository {
	return &getClientByCustomDomainRepository{
		database: database,
	}
}
