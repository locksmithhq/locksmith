package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/domain"
)

type updateSignupByClientIDRepository struct {
	database types.Database
}

// Execute implements contract.UpdateSignupByClientIDRepository.
func (r *updateSignupByClientIDRepository) Execute(ctx context.Context, signup domain.Signup) error {
	query := `
		UPDATE oauth_clients_signup
		SET
			custom_css = $2,
			custom_html = $3,
			input_variant = $4,
			layout = $5,
			show_social = $6,
			use_custom_html = $7,
			enabled = $8,
			background_color = $9,
			background_image = $10,
			background_type = $11,
			primary_color = $12,
			logo_url = $13,
			default_role_name = $14
		WHERE client_id = $1
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

func NewUpdateSignupByClientIDRepository(database types.Database) contract.UpdateSignupByClientIDRepository {
	return &updateSignupByClientIDRepository{database: database}
}
