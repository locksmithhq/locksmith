package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/contract"
	"github.com/booscaaa/locksmith/api/internal/oauth_clients_login/domain"
)

type getLoginByClientIDRepository struct {
	database types.Database
}

// Execute implements contract.GetLoginByClientIDRepository.
func (r *getLoginByClientIDRepository) Execute(ctx context.Context, clientID string) (domain.Login, error) {
	var login domain.Login

	query := `SELECT 
		id, client_id, custom_css, custom_html, input_variant, layout, 
		show_forgot_password, show_remember_me, show_sign_up, show_social, 
		use_custom_html, enabled, background_color, background_image, background_type,
		primary_color, logo_url,
		created_at, updated_at
	FROM oauth_clients_login 
	WHERE client_id = $1`

	err := r.database.QueryRowxContext(ctx, query, clientID).StructScan(&login)
	if err != nil {
		return domain.Login{}, stackerror.NewRepositoryError("GetLoginByClientIDRepository", err)
	}
	return login, nil
}

func NewGetLoginByClientIDRepository(database types.Database) contract.GetLoginByClientIDRepository {
	return &getLoginByClientIDRepository{database: database}
}
