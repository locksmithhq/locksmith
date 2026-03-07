package input

import (
	"github.com/locksmithhq/locksmith/api/internal/core/types/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/domain"
)

type Client struct {
	ProjectID    string   `json:"project_id"`
	Name         string   `json:"name"`
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURIs string   `json:"redirect_uris"`
	GrantTypes   string   `json:"grant_types"`
	CustomDomain string   `json:"custom_domain"`
	Roles        []string `json:"roles"`
}

func (c Client) ToDomain() domain.Client {
	return domain.NewClient(
		domain.WithProjectID(c.ProjectID),
		domain.WithName(c.Name),
		domain.WithClientID(c.ClientID),
		domain.WithClientSecret(c.ClientSecret),
		domain.WithRedirectURIs(c.RedirectURIs),
		domain.WithGrantTypes(c.GrantTypes),
		domain.WithCustomDomain(database.ParseNull(c.CustomDomain)),
	)
}
