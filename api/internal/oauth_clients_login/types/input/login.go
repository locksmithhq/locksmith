package input

import "github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/domain"

type Login struct {
	Layout             string `json:"layout"`
	InputVariant       string `json:"input_variant"`
	ShowSocial         bool   `json:"show_social"`
	ShowRememberMe     bool   `json:"show_remember_me"`
	ShowForgotPassword bool   `json:"show_forgot_password"`
	ShowSignUp         bool   `json:"show_sign_up"`
	UseCustomHTML      bool   `json:"use_custom_html"`
	Enabled            bool   `json:"enabled"`
	CustomHTML         string `json:"custom_html"`
	CustomCSS          string `json:"custom_css"`
	BackgroundColor    string `json:"background_color"`
	BackgroundImage    string `json:"background_image"`
	BackgroundType     string `json:"background_type"`
	PrimaryColor       string `json:"primary_color"`
	LogoURL            string `json:"logo_url"`
	FaviconURL         string `json:"favicon_url"`
}

func (in Login) ToLoginDomain() domain.Login {
	return domain.NewLogin(
		domain.WithLayout(in.Layout),
		domain.WithInputVariant(in.InputVariant),
		domain.WithShowSocial(in.ShowSocial),
		domain.WithShowRememberMe(in.ShowRememberMe),
		domain.WithShowForgotPassword(in.ShowForgotPassword),
		domain.WithShowSignUp(in.ShowSignUp),
		domain.WithUseCustomHTML(in.UseCustomHTML),
		domain.WithEnabled(in.Enabled),
		domain.WithCustomHTML(in.CustomHTML),
		domain.WithCustomCSS(in.CustomCSS),
		domain.WithBackgroundColor(in.BackgroundColor),
		domain.WithBackgroundImage(in.BackgroundImage),
		domain.WithBackgroundType(in.BackgroundType),
		domain.WithPrimaryColor(in.PrimaryColor),
		domain.WithLogoURL(in.LogoURL),
		domain.WithFaviconURL(in.FaviconURL),
	)
}
