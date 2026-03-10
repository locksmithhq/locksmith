package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getLoginByClientIDRepository struct {
	database types.Database
}

// Execute implements contract.GetLoginByClientIDRepository.
func (r *getLoginByClientIDRepository) Execute(ctx context.Context, clientID string) (domain.OAuthClientLogin, error) {
	var login domain.OAuthClientLogin

	query := `SELECT
		id, client_id, custom_css, custom_html, input_variant, layout,
		show_forgot_password, show_remember_me, show_sign_up, show_social,
		use_custom_html, enabled, background_color, background_image, background_type, primary_color, logo_url, favicon_url, created_at, updated_at
	FROM oauth_clients_login
	WHERE client_id = $1`

	err := r.database.GetContext(ctx, &login, query, clientID)
	if err != nil {
		return domain.OAuthClientLogin{}, stackerror.NewRepositoryError("GetLoginByClientIDRepository", err)
	}
	return login, nil
}

func NewGetLoginByClientIDRepository(database types.Database) contract.GetLoginByClientIDRepository {
	return &getLoginByClientIDRepository{database: database}
}
