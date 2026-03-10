package input

import (
	"github.com/locksmithhq/locksmith/api/internal/core/types/database"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/domain"
)

type Signup struct {
	Layout          string `json:"layout"`
	InputVariant    string `json:"input_variant"`
	ShowSocial      bool   `json:"show_social"`
	UseCustomHTML   bool   `json:"use_custom_html"`
	Enabled         bool   `json:"enabled"`
	CustomHTML      string `json:"custom_html"`
	CustomCSS       string `json:"custom_css"`
	BackgroundColor string `json:"background_color"`
	BackgroundImage string `json:"background_image"`
	BackgroundType  string `json:"background_type"`
	PrimaryColor    string `json:"primary_color"`
	LogoURL         string `json:"logo_url"`
	FaviconURL      string `json:"favicon_url"`
	DefaultRoleName string `json:"default_role_name"`
}

func (in Signup) ToSignupDomain() domain.Signup {
	return domain.Signup{
		Layout:          in.Layout,
		InputVariant:    in.InputVariant,
		ShowSocial:      in.ShowSocial,
		UseCustomHTML:   in.UseCustomHTML,
		Enabled:         in.Enabled,
		CustomHTML:      in.CustomHTML,
		CustomCSS:       in.CustomCSS,
		BackgroundColor: database.ParseNull(in.BackgroundColor),
		BackgroundImage: database.ParseNull(in.BackgroundImage),
		BackgroundType:  database.ParseNull(in.BackgroundType),
		PrimaryColor:    database.ParseNull(in.PrimaryColor),
		LogoURL:         database.ParseNull(in.LogoURL),
		FaviconURL:      database.ParseNull(in.FaviconURL),
		DefaultRoleName: in.DefaultRoleName,
	}
}
