package domain

type Client struct {
	ID           string `json:"id" db:"id"`
	ProjectID    string `json:"project_id" db:"project_id"`
	ClientID     string `json:"client_id" db:"client_id"`
	ClientSecret string `json:"client_secret" db:"client_secret"`
	RedirectURIs string `json:"redirect_uris" db:"redirect_uris"`
	GrantTypes   string `json:"grant_types" db:"grant_types"`
	Name         string `json:"name" db:"name"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
	DeletedAt    string `json:"deleted_at" db:"deleted_at"`
	Domain       string `json:"domain" db:"domain"`
}
