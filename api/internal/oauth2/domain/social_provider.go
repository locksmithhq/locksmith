package domain

type SocialProvider struct {
	ID           string `db:"id"`
	ClientID     string `db:"client_id"`
	Provider     string `db:"provider"`
	ClientKey    string `db:"client_key"`
	ClientSecret string `db:"client_secret"`
	Enabled      bool   `db:"enabled"`
	Scopes       string `db:"scopes"`
}

type SocialState struct {
	Nonce               string `db:"nonce"`
	ClientID            string `db:"client_id"`
	RedirectURI         string `db:"redirect_uri"`
	State               string `db:"state"`
	CodeChallenge       string `db:"code_challenge"`
	CodeChallengeMethod string `db:"code_challenge_method"`
	ExpiresAt           string `db:"expires_at"`
}

type AccountSocialProvider struct {
	ID             string `db:"id"`
	AccountID      string `db:"account_id"`
	Provider       string `db:"provider"`
	ProviderUserID string `db:"provider_user_id"`
	Email          string `db:"email"`
	CreatedAt      string `db:"created_at"`
}
