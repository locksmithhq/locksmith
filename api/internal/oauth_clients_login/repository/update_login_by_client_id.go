package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/domain"
)

type updateLoginByClientIDRepository struct {
	database types.Database
}

// Execute implements contract.UpdateLoginByClientIDRepository.
func (r *updateLoginByClientIDRepository) Execute(ctx context.Context, login domain.Login) error {
	query := `
		UPDATE oauth_clients_login
		SET
			custom_css = $2,
			custom_html = $3,
			input_variant = $4,
			layout = $5,
			show_forgot_password = $6,
			show_remember_me = $7,
			show_sign_up = $8,
			show_social = $9,
			use_custom_html = $10,
			enabled = $11,
			background_color = $12,
			background_image = $13,
			background_type = $14,
			primary_color = $15,
			logo_url = $16
		WHERE client_id = $1
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
	)
	if err != nil {
		return err
	}

	return nil
}

func NewUpdateLoginByClientIDRepository(database types.Database) contract.UpdateLoginByClientIDRepository {
	return &updateLoginByClientIDRepository{
		database: database,
	}
}
