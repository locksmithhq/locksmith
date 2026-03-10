package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth_clients_signup/domain"
)

type getSignupByClientIDRepository struct {
	database types.Database
}

// Execute implements contract.GetSignupByClientIDRepository.
func (r *getSignupByClientIDRepository) Execute(ctx context.Context, clientID string) (domain.Signup, error) {
	var signup domain.Signup

	query := `SELECT
		id, client_id, custom_css, custom_html, input_variant, layout,
		show_social, use_custom_html, enabled,
		background_color, background_image, background_type, primary_color, logo_url, favicon_url,
		default_role_name, created_at, updated_at
	FROM oauth_clients_signup
	WHERE client_id = $1`

	err := r.database.QueryRowxContext(ctx, query, clientID).StructScan(&signup)
	if err != nil {
		return domain.Signup{}, stackerror.NewRepositoryError("GetSignupByClientIDRepository", err)
	}
	return signup, nil
}

func NewGetSignupByClientIDRepository(database types.Database) contract.GetSignupByClientIDRepository {
	return &getSignupByClientIDRepository{database: database}
}
