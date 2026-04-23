package domain

type SocialProvider struct {
	ID           string `db:"id" json:"id"`
	ClientID     string `db:"client_id" json:"client_id"`
	Provider     string `db:"provider" json:"provider"`
	ClientKey    string `db:"client_key" json:"client_key"`
	ClientSecret string `db:"client_secret" json:"client_secret"`
	Enabled      bool   `db:"enabled" json:"enabled"`
	Scopes       string `db:"scopes" json:"scopes"`
	CreatedAt    string `db:"created_at" json:"created_at"`
	UpdatedAt    string `db:"updated_at" json:"updated_at"`
}
