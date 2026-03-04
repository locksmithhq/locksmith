package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
	"github.com/locksmithhq/locksmith/api/internal/oauth2/domain"
)

type createUserSessionRepository struct {
	database types.Database
}

// Execute implements contract.CreateUserSessionRepository.
func (r *createUserSessionRepository) Execute(ctx context.Context, entity domain.UserSession) (domain.UserSession, error) {
	var userSession domain.UserSession

	query := `INSERT INTO user_sessions (
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
		expires_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) 
	RETURNING 	
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
		created_at`

	err := r.database.QueryRowxContext(ctx,
		query,
		entity.AccountID,
		entity.ClientID,
		entity.JTI,
		entity.IPAddress,
		entity.UserAgent,
		entity.DeviceName,
		entity.DeviceID,
		entity.DeviceType,
		entity.Browser,
		entity.BrowserVersion,
		entity.OS,
		entity.OSVersion,
		entity.LocationCountry,
		entity.LocationRegion,
		entity.LocationCity,
		entity.ExpiresAt,
	).StructScan(&userSession)

	if err != nil {
		return domain.UserSession{}, stackerror.NewRepositoryError("CreateUserSessionRepository", err)
	}
	return userSession, nil
}

func NewCreateUserSessionRepository(database types.Database) contract.CreateUserSessionRepository {
	return &createUserSessionRepository{database: database}
}
