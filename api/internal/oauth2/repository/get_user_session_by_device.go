package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type getUserSessionByDeviceRepository struct {
	database types.Database
}

// Execute implements contract.GetUserSessionByDeviceRepository.
func (r *getUserSessionByDeviceRepository) Execute(ctx context.Context, accountID string, clientID string, deviceID string) (domain.UserSession, error) {
	var userSession domain.UserSession

	query := `SELECT 
				id, 
				account_id, 
				client_id, 
				jti, 
				ip_address, 
				user_agent, 
				device_name, 
				device_id, 
				device_type, 
				browser, 
				browser_version, 
				os, 
				os_version, 
				location_country, 
				location_region, 
				location_city, 
				expires_at, 
				revoked, 
				revoked_at, 
				revoked_reason, 
				last_activity, 
				created_at 
				FROM user_sessions 
				WHERE account_id = $1 AND client_id = $2 AND device_id = $3 AND revoked = false AND expires_at > NOW()
				LIMIT 1`

	err := r.database.GetContext(ctx, &userSession, query, accountID, clientID, deviceID)
	if err != nil {
		return domain.UserSession{}, stackerror.NewRepositoryError("GetUserSessionByDeviceRepository", err)
	}
	return userSession, nil
}

func NewGetUserSessionByDeviceRepository(database types.Database) contract.GetUserSessionByDeviceRepository {
	return &getUserSessionByDeviceRepository{database: database}
}
