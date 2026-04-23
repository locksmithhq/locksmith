package input

type SocialProvider struct {
	ClientKey    string `json:"client_key"`
	ClientSecret string `json:"client_secret"`
	Enabled      bool   `json:"enabled"`
	Scopes       string `json:"scopes"`
}
