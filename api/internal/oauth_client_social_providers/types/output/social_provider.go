package output

import "github.com/locksmithhq/locksmith/api/internal/oauth_client_social_providers/domain"

type SocialProvider struct {
	ID           string `json:"id"`
	ClientID     string `json:"client_id"`
	Provider     string `json:"provider"`
	ClientKey    string `json:"client_key"`
	ClientSecret string `json:"client_secret"`
	Enabled      bool   `json:"enabled"`
	Scopes       string `json:"scopes"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func NewSocialProviderFromDomain(p domain.SocialProvider) SocialProvider {
	return SocialProvider{
		ID:           p.ID,
		ClientID:     p.ClientID,
		Provider:     p.Provider,
		ClientKey:    p.ClientKey,
		ClientSecret: "****",
		Enabled:      p.Enabled,
		Scopes:       p.Scopes,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func NewSocialProvidersFromDomain(providers []domain.SocialProvider) []SocialProvider {
	out := make([]SocialProvider, len(providers))
	for i, p := range providers {
		out[i] = NewSocialProviderFromDomain(p)
	}
	return out
}
