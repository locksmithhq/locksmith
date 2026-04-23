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
	RequirePKCE  bool          `json:"require_pkce"`
}

func maskSecret(secret string) string {
	if len(secret) <= 4 {
		return "****"
	}
	return "****" + secret[len(secret)-4:]
}

// NewClientFromDomain maps a domain client to the output type with the
// client_secret masked. Use NewCreatedClientFromDomain on creation so the
// full secret is returned exactly once.
func NewClientFromDomain(client domain.Client) Client {
	return Client{
		ID:           client.ID,
		ProjectID:    client.ProjectID,
		ClientID:     client.ClientID,
		ClientSecret: maskSecret(client.ClientSecret),
		RedirectURIs: client.RedirectURIs,
		GrantTypes:   client.GrantTypes,
		Name:         client.Name,
		CreatedAt:    client.CreatedAt,
		UpdatedAt:    client.UpdatedAt,
		CustomDomain: client.CustomDomain,
		RequirePKCE:  client.RequirePKCE,
	}
}

// NewCreatedClientFromDomain is identical to NewClientFromDomain but returns
// the full client_secret. Must only be called from the create flow.
func NewCreatedClientFromDomain(client domain.Client) Client {
	out := NewClientFromDomain(client)
	out.ClientSecret = client.ClientSecret
	return out
}

func NewClientsFromDomain(c []domain.Client) []Client {
	var out []Client = make([]Client, len(c))

	for i, client := range c {
		out[i] = NewClientFromDomain(client)
	}

	return out
}
