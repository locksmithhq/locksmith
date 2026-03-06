package repository

import (
	"context"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/locksmithhq/locksmith/api/internal/core/types/stackerror"
	"github.com/locksmithhq/locksmith/api/internal/session/contract"
	"github.com/locksmithhq/locksmith/api/internal/session/domain"
)

type fetchSessionsByAccountIDRepository struct {
	database types.Database
}

func (r *fetchSessionsByAccountIDRepository) Execute(ctx context.Context, projectID, accountID string, page, limit int) ([]domain.Session, error) {
	var sessions []domain.Session
	offset := (page - 1) * limit

	query := `
		SELECT
			us.id,
			us.account_id,
			us.client_id,
			COALESCE(us.ip_address::text, '') AS ip_address,
			COALESCE(us.device_type, '') AS device_type,
			COALESCE(us.browser, '') AS browser,
			COALESCE(us.browser_version, '') AS browser_version,
			COALESCE(us.os, '') AS os,
			COALESCE(us.os_version, '') AS os_version,
			COALESCE(us.location_country, '') AS location_country,
			COALESCE(us.location_region, '') AS location_region,
			COALESCE(us.location_city, '') AS location_city,
			us.revoked,
			COALESCE(us.revoked_reason, '') AS revoked_reason,
			to_char(us.expires_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS expires_at,
			COALESCE(to_char(us.last_activity AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"'), '') AS last_activity,
			to_char(us.created_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS created_at,
			a.name AS account_name,
			a.email AS account_email,
			oc.name AS client_name
		FROM user_sessions us
		JOIN accounts a ON a.id = us.account_id
		JOIN oauth_clients oc ON oc.id = us.client_id
		WHERE oc.project_id = $1 AND us.account_id = $2
		ORDER BY us.created_at DESC
		LIMIT $3 OFFSET $4
	`

	err := r.database.SelectContext(ctx, &sessions, query, projectID, accountID, limit, offset)
	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchSessionsByAccountIDRepository", err)
	}

	return sessions, nil
}

func NewFetchSessionsByAccountIDRepository(database types.Database) contract.FetchSessionsByAccountIDRepository {
	return &fetchSessionsByAccountIDRepository{database: database}
}
