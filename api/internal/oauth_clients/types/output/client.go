package output

import (
	"github.com/locksmithhq/locksmith/api/internal/core/types/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients/domain"
)

type Client struct {
	ID           string        `json:"id"`
	ProjectID    string        `json:"project_id"`
	ClientID     string        `json:"client_id"`
	ClientSecret string        `json:"client_secret"`
	RedirectURIs string        `json:"redirect_uris"`
	GrantTypes   string        `json:"grant_types"`
	Name         string        `json:"name"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
	Roles        []string      `json:"roles"`
	CustomDomain database.Null `json:"custom_domain"`
}

func NewClientFromDomain(client domain.Client) Client {
	return Client{
		ID:           client.ID,
		ProjectID:    client.ProjectID,
		ClientID:     client.ClientID,
		ClientSecret: client.ClientSecret,
		RedirectURIs: client.RedirectURIs,
		GrantTypes:   client.GrantTypes,
		Name:         client.Name,
		CreatedAt:    client.CreatedAt,
		UpdatedAt:    client.UpdatedAt,
		CustomDomain: client.CustomDomain,
	}
}

func NewClientsFromDomain(c []domain.Client) []Client {
	var out []Client = make([]Client, len(c))

	for i, client := range c {
		out[i] = NewClientFromDomain(client)
	}

	return out
}
