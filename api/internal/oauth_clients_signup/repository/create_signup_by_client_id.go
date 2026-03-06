package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/domain"
)

type createSignupByClientIDRepository struct {
	database types.Database
}

// Execute implements contract.CreateSignupByClientIDRepository.
func (r *createSignupByClientIDRepository) Execute(ctx context.Context, signup domain.Signup) error {
	query := `
		INSERT INTO oauth_clients_signup (
			client_id,
			custom_css,
			custom_html,
			input_variant,
			layout,
			show_social,
			use_custom_html,
			enabled,
			background_color,
			background_image,
			background_type,
			primary_color,
			logo_url,
			default_role_name
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`

	_, err := r.database.ExecContext(
		ctx,
		query,
		signup.ClientID,
		signup.CustomCSS,
		signup.CustomHTML,
		signup.InputVariant,
		signup.Layout,
		signup.ShowSocial,
		signup.UseCustomHTML,
		signup.Enabled,
		signup.BackgroundColor,
		signup.BackgroundImage,
		signup.BackgroundType,
		signup.PrimaryColor,
		signup.LogoURL,
		signup.DefaultRoleName,
	)
	if err != nil {
		return err
	}

	return nil
}

func NewCreateSignupByClientIDRepository(database types.Database) contract.CreateSignupByClientIDRepository {
	return &createSignupByClientIDRepository{database: database}
}
