package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_login/domain"
)

type createLoginByClientIDRepository struct {
	database types.Database
}

// Execute implements contract.CreateLoginByClientIDRepository.
func (r *createLoginByClientIDRepository) Execute(ctx context.Context, login domain.Login) error {
	query := `
		INSERT INTO oauth_clients_login (
			client_id,
			custom_css,
			custom_html,
			input_variant,
			layout,
			show_forgot_password,
			show_remember_me,
			show_sign_up,
			show_social,
			use_custom_html,
			enabled,
			background_color,
			background_image,
			background_type,
			primary_color,
			logo_url,
			favicon_url
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		ON CONFLICT (client_id) DO NOTHING
	`

	_, err := r.database.ExecContext(
		ctx,
		query,
		login.ClientID,
		login.CustomCSS,
		login.CustomHTML,
		login.InputVariant,
		login.Layout,
		login.ShowForgotPassword,
		login.ShowRememberMe,
		login.ShowSignUp,
		login.ShowSocial,
		login.UseCustomHTML,
		login.Enabled,
		login.BackgroundColor,
		login.BackgroundImage,
		login.BackgroundType,
		login.PrimaryColor,
		login.LogoURL,
		login.FaviconURL,
	)
	if err != nil {
		return err
	}

	return nil
}

func NewCreateLoginByClientIDRepository(database types.Database) contract.CreateLoginByClientIDRepository {
	return &createLoginByClientIDRepository{
		database: database,
	}
}
