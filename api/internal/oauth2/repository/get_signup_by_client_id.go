package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getSignupByClientIDRepository struct {
	database types.Database
}

// Execute implements contract.GetSignupByClientIDRepository.
func (r *getSignupByClientIDRepository) Execute(ctx context.Context, clientID string) (domain.OAuthClientSignup, error) {
	var signup domain.OAuthClientSignup

	query := `SELECT
		id, client_id, custom_css, custom_html, input_variant, layout,
		show_social, use_custom_html, enabled,
		background_color, background_image, background_type, primary_color, logo_url,
		created_at, updated_at
	FROM oauth_clients_signup
	WHERE client_id = $1`

	err := r.database.GetContext(ctx, &signup, query, clientID)
	if err != nil {
		return domain.OAuthClientSignup{}, stackerror.NewRepositoryError("GetSignupByClientIDRepository", err)
	}
	return signup, nil
}

func NewGetSignupByClientIDRepository(database types.Database) contract.GetSignupByClientIDRepository {
	return &getSignupByClientIDRepository{database: database}
}
