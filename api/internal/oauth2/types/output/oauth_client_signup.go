package output

import "github.com/locksmithhq/locksmith/api/internal/core/types/database"

type OAuthClientSignup struct {
	ID              string        `json:"id"`
	ClientID        string        `json:"client_id"`
	CustomCSS       string        `json:"custom_css"`
	CustomHTML      string        `json:"custom_html"`
	InputVariant    string        `json:"input_variant"`
	Layout          string        `json:"layout"`
	ShowSocial      bool          `json:"show_social"`
	UseCustomHTML   bool          `json:"use_custom_html"`
	Enabled         bool          `json:"enabled"`
	BackgroundColor database.Null `json:"background_color"`
	BackgroundImage database.Null `json:"background_image"`
	BackgroundType  database.Null `json:"background_type"`
	PrimaryColor    database.Null `json:"primary_color"`
	LogoURL         database.Null `json:"logo_url"`
	FaviconURL      database.Null `json:"favicon_url"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
}
