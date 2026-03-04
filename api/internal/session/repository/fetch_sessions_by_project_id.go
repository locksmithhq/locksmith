package repository

import (
	"context"
	"fmt"

	"github.com/booscaaa/initializers/postgres/types"
	"github.com/booscaaa/locksmith/api/internal/core/types/stackerror"
	"github.com/booscaaa/locksmith/api/internal/session/contract"
	"github.com/booscaaa/locksmith/api/internal/session/domain"
)

type fetchSessionsByProjectIDRepository struct {
	database types.Database
}

func (r *fetchSessionsByProjectIDRepository) Execute(ctx context.Context, projectID string, page, limit int, search string) ([]domain.Session, error) {
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
			COALESCE(us.location_city, '') AS location_city,
			us.revoked,
			COALESCE(us.revoked_reason, '') AS revoked_reason,
			us.expires_at::text AS expires_at,
			us.last_activity::text AS last_activity,
			us.created_at::text AS created_at,
			a.name AS account_name,
			a.email AS account_email,
			oc.name AS client_name
		FROM user_sessions us
		JOIN accounts a ON a.id = us.account_id
		JOIN oauth_clients oc ON oc.id = us.client_id
		WHERE oc.project_id = $1
	`

	args := []interface{}{projectID}
	argIdx := 2

	if search != "" {
		query += fmt.Sprintf(
			` AND (a.name ILIKE $%d OR a.email ILIKE $%d OR COALESCE(us.ip_address::text, '') ILIKE $%d OR COALESCE(us.browser, '') ILIKE $%d)`,
			argIdx, argIdx, argIdx, argIdx,
		)
		args = append(args, "%"+search+"%")
		argIdx++
	}

	query += fmt.Sprintf(` ORDER BY us.created_at DESC LIMIT $%d OFFSET $%d`, argIdx, argIdx+1)
	args = append(args, limit, offset)

	err := r.database.SelectContext(ctx, &sessions, query, args...)
	if err != nil {
		return nil, stackerror.NewRepositoryError("FetchSessionsByProjectIDRepository", err)
	}

	return sessions, nil
}

func NewFetchSessionsByProjectIDRepository(database types.Database) contract.FetchSessionsByProjectIDRepository {
	return &fetchSessionsByProjectIDRepository{database: database}
}
