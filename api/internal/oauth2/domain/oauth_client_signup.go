package domain

import "github.com/locksmithhq/locksmith/api/internal/core/types/database"

type OAuthClientSignup struct {
	ID              string        `json:"id" db:"id" paginate:"id"`
	ClientID        string        `json:"client_id" db:"client_id" paginate:"client_id"`
	CustomCSS       string        `json:"custom_css" db:"custom_css" paginate:"custom_css"`
	CustomHTML      string        `json:"custom_html" db:"custom_html" paginate:"custom_html"`
	InputVariant    string        `json:"input_variant" db:"input_variant" paginate:"input_variant"`
	Layout          string        `json:"layout" db:"layout" paginate:"layout"`
	ShowSocial      bool          `json:"show_social" db:"show_social" paginate:"show_social"`
	UseCustomHTML   bool          `json:"use_custom_html" db:"use_custom_html" paginate:"use_custom_html"`
	Enabled         bool          `json:"enabled" db:"enabled" paginate:"enabled"`
	BackgroundColor database.Null `json:"background_color" db:"background_color" paginate:"background_color"`
	BackgroundImage database.Null `json:"background_image" db:"background_image" paginate:"background_image"`
	BackgroundType  database.Null `json:"background_type" db:"background_type" paginate:"background_type"`
	PrimaryColor    database.Null `json:"primary_color" db:"primary_color" paginate:"primary_color"`
	LogoURL         database.Null `json:"logo_url" db:"logo_url" paginate:"logo_url"`
	CreatedAt       string        `json:"created_at" db:"created_at" paginate:"created_at"`
	UpdatedAt       string        `json:"updated_at" db:"updated_at" paginate:"updated_at"`
}
