package domain

import "github.com/locksmithhq/locksmith/api/internal/core/types/database"

type Client struct {
	ID           string        `json:"id" db:"id" paginate:"id"`
	ProjectID    string        `json:"project_id" db:"project_id" paginate:"project_id"`
	ClientID     string        `json:"client_id" db:"client_id" paginate:"client_id"`
	ClientSecret string        `json:"client_secret" db:"client_secret" paginate:"client_secret"`
	RedirectURIs string        `json:"redirect_uris" db:"redirect_uris" paginate:"redirect_uris"`
	GrantTypes   string        `json:"grant_types" db:"grant_types" paginate:"grant_types"`
	Name         string        `json:"name" db:"name" paginate:"name"`
	CreatedAt    string        `json:"created_at" db:"created_at" paginate:"created_at"`
	UpdatedAt    string        `json:"updated_at" db:"updated_at" paginate:"updated_at"`
	DeletedAt    string        `json:"deleted_at" db:"deleted_at" paginate:"deleted_at"`
	CustomDomain database.Null `json:"custom_domain" db:"custom_domain" paginate:"custom_domain"`
	RequirePKCE  bool         `json:"require_pkce" db:"require_pkce" paginate:"require_pkce"`
}

type ClientOption func(*Client)

func WithID(id string) ClientOption {
	return func(c *Client) {
		c.ID = id
	}
}

func WithProjectID(projectID string) ClientOption {
	return func(c *Client) {
		c.ProjectID = projectID
	}
}

func WithClientID(clientID string) ClientOption {
	return func(c *Client) {
		c.ClientID = clientID
	}
}

func WithClientSecret(clientSecret string) ClientOption {
	return func(c *Client) {
		c.ClientSecret = clientSecret
	}
}

func WithRedirectURIs(redirectURIs string) ClientOption {
	return func(c *Client) {
		c.RedirectURIs = redirectURIs
	}
}

func WithGrantTypes(grantTypes string) ClientOption {
	return func(c *Client) {
		c.GrantTypes = grantTypes
	}
}

func WithName(name string) ClientOption {
	return func(c *Client) {
		c.Name = name
	}
}

func WithCreatedAt(createdAt string) ClientOption {
	return func(c *Client) {
		c.CreatedAt = createdAt
	}
}

func WithUpdatedAt(updatedAt string) ClientOption {
	return func(c *Client) {
		c.UpdatedAt = updatedAt
	}
}

func WithDeletedAt(deletedAt string) ClientOption {
	return func(c *Client) {
		c.DeletedAt = deletedAt
	}
}

func WithCustomDomain(customDomain database.Null) ClientOption {
	return func(c *Client) {
		c.CustomDomain = customDomain
	}
}

func WithRequirePKCE(requirePKCE bool) ClientOption {
	return func(c *Client) {
		c.RequirePKCE = requirePKCE
	}
}

func NewClient(options ...ClientOption) Client {
	client := Client{}

	for _, option := range options {
		option(&client)
	}

	return client
}
